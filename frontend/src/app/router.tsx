import { createBrowserRouter } from "react-router";
import { HomePage } from "@/page/home";
import { Suspense } from "react";
import { AuthLayout, SigninPage, SignupPage } from "@/page/auth";

export const router = createBrowserRouter([
  {
    path: "/",
    children: [
      {
        index: true,
        element: (
          <Suspense fallback='loading'>
            <HomePage />
          </Suspense>
        ),
      },
      {
        path: "auth",
        Component: AuthLayout,
        children: [
          {
            path: "login",
            element: (
              <Suspense fallback='loading'>
                <SigninPage />
              </Suspense>
            ),
          },
          {
            path: "register",
            element: (
              <Suspense fallback='loading'>
                <SignupPage />
              </Suspense>
            ),
          },
        ],
      },
    ],
  },
]);
