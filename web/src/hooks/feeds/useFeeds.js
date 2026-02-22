import { useQuery } from "@tanstack/react-query";
import { getFeeds } from "../../api/feeds";

export function useFeeds() {
  return useQuery({
    queryKey: ["feeds"],
    queryFn: getFeeds,
  });
}