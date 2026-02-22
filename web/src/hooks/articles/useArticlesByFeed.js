import { useQuery } from "@tanstack/react-query";
import { getArticlesByFeed } from "../../api/articles";

export function useArticlesByFeed(feedId) {
  return useQuery({
    queryKey: ["articles", "byFeed", feedId],
    queryFn: () => getArticlesByFeed(feedId),
    enabled: !!feedId,
  });
}