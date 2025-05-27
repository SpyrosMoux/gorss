import { createBrowserRouter, RouterProvider } from "react-router";
import { NewsFeed } from "../pages/newsFeed/NewsFeed";
import { MainLayout } from "../layouts/MainLayout";

const router = createBrowserRouter([
  {
    Component: MainLayout,
    children: [{ index: true, Component: NewsFeed }],
  },
]);

export const Router = () => {
  return <RouterProvider router={router} />;
};
