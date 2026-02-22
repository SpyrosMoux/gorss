import axiosApp from "./http";

export const getLatestArticles = async () => {
  const response = await axiosApp.get("/articles/latest");
  return response.data;
};