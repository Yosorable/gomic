/* @refresh reload */
import { render } from "solid-js/web";
import { lazy, Suspense } from "solid-js";
import { HashRouter } from "@solidjs/router";
import { MetaProvider } from "@solidjs/meta";
import "./index.css";

const routes = [
  {
    path: "/",
    component: lazy(() => import("./views/Home")),
  },
  {
    path: "/authors",
    component: lazy(() => import("./views/Authors")),
  },
  {
    path: "/author/:name",
    component: lazy(() => import("./views/Author")),
  },
  {
    path: "/viewer/:name",
    component: lazy(() => import("./views/Viewer")),
  },
  {
    path: "/syncing",
    component: lazy(() => import("./views/SyncingFiles")),
  },
];

render(
  () => (
    <HashRouter
      root={(props) => (
        <MetaProvider>
          <Suspense>{props.children}</Suspense>
        </MetaProvider>
      )}
    >
      {routes}
    </HashRouter>
  ),
  document.querySelector("#root")!
);
