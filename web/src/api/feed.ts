import type { GetAllFeedsResponse } from "../types/feed";
import axiosApp from "./axios-instance";

export const getFeed = async (): Promise<GetAllFeedsResponse> => {
  const response = await axiosApp.get("/feeds");
  return response.data;
};
