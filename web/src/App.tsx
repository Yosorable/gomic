import { lazy, Show, Suspense } from "solid-js";
import { HashRouter, Navigate } from "@solidjs/router";
import "./index.scss";
import Layout from "./views/components/Layout";

const routes = [
  {
    path: "/home",
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
  {
    path: "/",
    component: () => <Navigate href="/home" />,
  },
];

export default function App() {
  return (
    <HashRouter
      root={(props) => (
        <Show
          when={!props.location.pathname.startsWith("/viewer")}
          fallback={props.children}
        >
          <Layout>
            <Suspense>{props.children}</Suspense>
          </Layout>
        </Show>
      )}
    >
      {routes}
    </HashRouter>
  );
}
