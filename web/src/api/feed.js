import axiosApp from "./axios-instance";

export const getFeeds = async () => {
  const response = await axiosApp.get("/feeds");
  return response.data;
};