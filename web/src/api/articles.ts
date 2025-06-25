import type { ArticlesResponse } from "../types/articles";
import axiosApp from "./axios-instance";

export const getLatestArticles = async (): Promise<ArticlesResponse> => {
  const response = await axiosApp.get("/articles/latest");
  return response.data;
};
