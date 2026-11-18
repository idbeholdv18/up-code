import { FC } from "react";
import { Outlet } from "react-router";

export interface SigninPageProps {}

export const AuthLayout: FC<SigninPageProps> = (props) => {
  return (
    <div>
      <h1>Auth layout</h1>
      <Outlet />
    </div>
  );
};
