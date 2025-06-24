import { useQuery } from "@tanstack/react-query";
import { getFeed } from "../../api/feed";

const useFeed = () => {
  return useQuery({
    queryKey: ["get-feed"],
    queryFn: getFeed,
  });
};

export default useFeed;
