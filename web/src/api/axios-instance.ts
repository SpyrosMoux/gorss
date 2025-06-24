import axios from "axios";

const axiosApp = axios.create({
  baseURL: "https://gorss-api.spyrosmoux.com/api",
});

export default axiosApp;
