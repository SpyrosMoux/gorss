import axiosApp from "./http";

export const getFeeds = async () => {
  const { data } = await axiosApp.get("/feeds");
  return Array.isArray(data?.feeds) ? data.feeds : [];
};