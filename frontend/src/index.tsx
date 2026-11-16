import "react";
import { createRoot } from "react-dom/client";

document.body.innerHTML = '<div id="app"></div>';

// Render your React component instead
const appElement = document.getElementById("app");
if (!appElement) {
  throw new Error("element with id app doesn't exist");
}
const root = createRoot(appElement);
root.render(<h1>Up code</h1>);
