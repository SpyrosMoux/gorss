import { Outlet } from "react-router-dom";
import Header from "./components/Header";

export function MainLayout() {
  return (
    <>
      <Header />
      <main className="max-w-2xl mx-auto px-4 py-6">
        <Outlet />
      </main>
    </>
  );
}