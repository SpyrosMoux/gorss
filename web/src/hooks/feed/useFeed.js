import { useQuery } from "@tanstack/react-query";
import { getFeeds } from "../../api/feed";

const useFeeds = () => {
  return useQuery({
    queryKey: ["feeds"],
    queryFn: getFeeds,
  });
};

export default useFeeds;
