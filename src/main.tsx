import { StrictMode } from "react";
import ReactDOM from "react-dom/client";
import { RouterProvider, Router } from "@tanstack/react-router";
import { rootRoute } from "./routes/root/root-route";
import { indexRoute } from "./routes/index/index-route";

const routeTree = rootRoute.addChildren([indexRoute]);

const router = new Router({ routeTree });

declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router;
  }
}

const rootElement = document.getElementById("root")!;
if (!rootElement.innerHTML) {
  const root = ReactDOM.createRoot(rootElement);
  root.render(
    <StrictMode>
      <RouterProvider router={router} />
    </StrictMode>,
  );
}
