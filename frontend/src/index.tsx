import "react";
import { createRoot } from "react-dom/client";
import { RouterProvider } from "react-router";
import { router } from "./app/router";
import { Provider } from "react-redux";
import { store } from "@/shared/store/store";
import "./app/index.css";

document.body.innerHTML = '<div id="app"></div>';

const appElement = document.getElementById("app");
if (!appElement) {
  throw new Error("element with id app doesn't exist");
}
const root = createRoot(appElement);
root.render(
  <Provider store={store}>
    <RouterProvider router={router} />
  </Provider>,
);
