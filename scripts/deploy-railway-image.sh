#!/usr/bin/env bash

set -euo pipefail
IFS=$'\n\t'

readonly GRAPHQL_ENDPOINT="https://backboard.railway.com/graphql/v2"
readonly AUTH_MODE_BEARER="bearer"
readonly AUTH_MODE_PROJECT="project"

railway_token=""
railway_project_id=""
railway_environment_name=""
railway_environment_id=""
railway_service_id=""
image_repository=""
image_tag=""
token_type="$AUTH_MODE_BEARER"

usage() {
  printf 'Usage: %s --railway-token <token> --railway-project-id <id> (--railway-environment-id <id> | --railway-environment-name <name>) --railway-service-id <id> --image-repository <repo> --image-tag <tag> [--token-type bearer|project]\n' "$0" >&2
}

log() {
  printf '%s\n' "$*"
}

die() {
  printf 'Error: %s\n' "$*" >&2
  exit 1
}

require_tools() {
  command -v curl >/dev/null 2>&1 || die "curl is required but not installed"
  command -v jq >/dev/null 2>&1 || die "jq is required but not installed"
}

parse_args() {
  while [[ $# -gt 0 ]]; do
    case "$1" in
      --railway-token)
        [[ $# -ge 2 ]] || die "Missing value for --railway-token"
        railway_token="$2"
        shift 2
        ;;
      --railway-project-id)
        [[ $# -ge 2 ]] || die "Missing value for --railway-project-id"
        railway_project_id="$2"
        shift 2
        ;;
      --railway-environment-name)
        [[ $# -ge 2 ]] || die "Missing value for --railway-environment-name"
        railway_environment_name="$2"
        shift 2
        ;;
      --railway-environment-id)
        [[ $# -ge 2 ]] || die "Missing value for --railway-environment-id"
        railway_environment_id="$2"
        shift 2
        ;;
      --railway-service-id)
        [[ $# -ge 2 ]] || die "Missing value for --railway-service-id"
        railway_service_id="$2"
        shift 2
        ;;
      --image-repository)
        [[ $# -ge 2 ]] || die "Missing value for --image-repository"
        image_repository="$2"
        shift 2
        ;;
      --image-tag)
        [[ $# -ge 2 ]] || die "Missing value for --image-tag"
        image_tag="$2"
        shift 2
        ;;
      --token-type)
        [[ $# -ge 2 ]] || die "Missing value for --token-type"
        token_type="$2"
        shift 2
        ;;
      -h|--help)
        usage
        exit 0
        ;;
      *)
        usage
        die "Unknown argument: $1"
        ;;
    esac
  done
}

validate_args() {
  local required
  for required in railway_token railway_project_id railway_service_id image_repository image_tag; do
    [[ -n "${!required}" ]] || die "Missing required argument: $required"
  done

  if [[ -z "$railway_environment_id" && -z "$railway_environment_name" ]]; then
    die "Provide either --railway-environment-id or --railway-environment-name"
  fi

  if [[ "$token_type" != "$AUTH_MODE_BEARER" && "$token_type" != "$AUTH_MODE_PROJECT" ]]; then
    die "Invalid --token-type. Allowed values: bearer, project"
  fi
}

graphql_request() {
  local query="$1"
  local variables_json="$2"
  local payload response
  local -a headers

  payload="$(jq -cn --arg query "$query" --argjson variables "$variables_json" '{query: $query, variables: $variables}')"

  headers=(--header "Content-Type: application/json")
  if [[ "$token_type" == "$AUTH_MODE_PROJECT" ]]; then
    headers+=(--header "Project-Access-Token: $railway_token")
  else
    headers+=(--header "Authorization: Bearer $railway_token")
  fi

  response="$(curl --silent --show-error --fail \
    --request POST \
    --url "$GRAPHQL_ENDPOINT" \
    "${headers[@]}" \
    --data "$payload")"

  if jq -e '.errors and (.errors | length > 0)' >/dev/null <<<"$response"; then
    printf 'Railway GraphQL request failed:\n' >&2
    jq '.errors' <<<"$response" >&2
    exit 1
  fi

  printf '%s\n' "$response"
}

resolve_environment_id() {
  if [[ -n "$railway_environment_id" ]]; then
    printf '%s\n' "$railway_environment_id"
    return
  fi

  local query variables response resolved_id

  query='query Environments($projectId: String!) {
  environments(projectId: $projectId) {
    edges {
      node {
        id
        name
      }
    }
  }
}'

  variables="$(jq -cn --arg projectId "$railway_project_id" '{projectId: $projectId}')"
  response="$(graphql_request "$query" "$variables")"
  resolved_id="$(jq -r --arg envName "$railway_environment_name" '[.data.environments.edges[].node | select(.name == $envName) | .id][0] // empty' <<<"$response")"

  [[ -n "$resolved_id" ]] || die "Could not find environment named '$railway_environment_name' in project '$railway_project_id'"
  printf '%s\n' "$resolved_id"
}

update_service_image() {
  local environment_id="$1"
  local image_reference="$2"
  local mutation variables

  mutation='mutation UpdateServiceSource($serviceId: String!, $environmentId: String!, $image: String!) {
  serviceInstanceUpdate(
    serviceId: $serviceId
    environmentId: $environmentId
    input: { source: { image: $image } }
  )
}'

  variables="$(jq -cn \
    --arg serviceId "$railway_service_id" \
    --arg environmentId "$environment_id" \
    --arg image "$image_reference" \
    '{serviceId: $serviceId, environmentId: $environmentId, image: $image}')"

  graphql_request "$mutation" "$variables" >/dev/null
}

trigger_deploy() {
  local environment_id="$1"
  local mutation variables response deployment_id

  mutation='mutation DeployService($serviceId: String!, $environmentId: String!) {
  serviceInstanceDeployV2(serviceId: $serviceId, environmentId: $environmentId)
}'

  variables="$(jq -cn --arg serviceId "$railway_service_id" --arg environmentId "$environment_id" '{serviceId: $serviceId, environmentId: $environmentId}')"
  response="$(graphql_request "$mutation" "$variables")"
  deployment_id="$(jq -r '.data.serviceInstanceDeployV2 // empty' <<<"$response")"

  [[ -n "$deployment_id" ]] || die "Railway deploy did not return a deployment id"
  printf '%s\n' "$deployment_id"
}

main() {
  parse_args "$@"
  validate_args
  require_tools

  local environment_id image_reference deployment_id

  image_reference="${image_repository}:${image_tag}"
  log "Preparing Railway deploy for $image_reference"

  environment_id="$(resolve_environment_id)"
  update_service_image "$environment_id" "$image_reference"
  log "Updated Railway service source image to $image_reference"

  deployment_id="$(trigger_deploy "$environment_id")"
  log "Triggered Railway deployment: $deployment_id"
}

main "$@"
