import { Outlet } from "react-router";
import { Navbar } from "./components/Navbar";
import { Container } from "@mui/material";

export const MainLayout = () => {
  return (
    <>
      <Navbar />
      <Container maxWidth="md">
        <Outlet />
      </Container>
    </>
  );
};
