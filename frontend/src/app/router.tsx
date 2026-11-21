import { createBrowserRouter } from "react-router";
import { HomePage } from "@/page/home";
import { Suspense } from "react";
import { AuthLayout, SigninPage, SignupPage } from "@/page/auth";
import { RootLayout } from "@/page/root";
import { DrawerProvider } from "@/shared/ui/drawer";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <RootLayout />,
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
