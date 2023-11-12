import "@/index.css";
import { Outlet } from "@tanstack/react-router";

const Root = () => {
  return (
    <div>
      <Outlet />
    </div>
  );
};

export default Root;
