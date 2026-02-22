import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { MainLayout } from "../layouts/MainLayout";
import { Home } from "../../pages/home/Home";
import { Feeds } from "../../pages/feeds/Feeds";
import { FeedArticles } from "../../pages/feed-articles/FeedArticles";

const router = createBrowserRouter([
  {
    path: "/",
    Component: MainLayout,
    children: [
      { index: true, Component: Home },
      { path: "feeds", Component: Feeds },
      { path: "feeds/:feedId", Component: FeedArticles },
    ],
  },
]);

export function Router() {
  return <RouterProvider router={router} />;
}