import { useQuery } from "@tanstack/react-query";
import { getLatestArticles } from "../../api/articles";

export function useLatestArticles() {
  return useQuery({
    queryKey: ["articles", "latest"],
    queryFn: getLatestArticles,
  });
}