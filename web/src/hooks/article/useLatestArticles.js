import { useQuery } from "@tanstack/react-query";
import { getLatestArticles } from "../../api/articles";

const useLatestArticles = () => {
  return useQuery({
    queryKey: ["get-latest-articles"],
    queryFn: getLatestArticles,
  });
};

export default useLatestArticles;