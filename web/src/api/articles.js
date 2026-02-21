import axiosApp from "./axios-instance";

export const getLatestArticles = async () => {
  const response = await axiosApp.get("/articles/latest");
  return response.data;
};