import axiosApp from "./http";

export const getLatestArticles = async () => {
  const { data } = await axiosApp.get("/articles/latest");
  return Array.isArray(data?.articles) ? data.articles : [];
};