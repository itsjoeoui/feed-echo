import { Route } from "@tanstack/react-router";
import Index from ".";
import { rootRoute } from "../root/root-route";

export const indexRoute = new Route({
  getParentRoute: () => rootRoute,
  path: "/",
  component: Index,
});
