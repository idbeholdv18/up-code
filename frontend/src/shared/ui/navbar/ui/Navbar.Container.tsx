import clsx from "clsx";
import { FC, PropsWithChildren } from "react";

export type NavbarContainerProps = PropsWithChildren<{
  className?: string;
}>;

export const NavbarContainer: FC<NavbarContainerProps> = (props) => {
  return <div className={clsx(props.className, "items-center")}>{props.children}</div>;
};
