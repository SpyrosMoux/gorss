import { Outlet } from "react-router";
import { Navbar } from "./components/Navbar";

export const MainLayout = () => {
  return (
    <>
      <Navbar />
      <Outlet />
    </>
  );
};
