import { Outlet } from "react-router";
import Navigation from "../Navigation/Navigation";

export default function Layout() {
  return (
    <div className="w-[95vw] grid grid-rows-[0.5fr_5fr] min-h-screen">
      <Navigation />
      <div className="mt-5">
        <Outlet />
      </div>
    </div>
  );
}
