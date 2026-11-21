import clsx from "clsx";
import { FC, PropsWithChildren } from "react";

export type NavbarBaseProps = PropsWithChildren<{
  className?: string;
}>;

export const NavbarBase: FC<NavbarBaseProps> = (props) => {
  return (
    <div className={clsx("w-full fixed top-0 z-50", props.className)}>
      {props.children}
    </div>
  );
};
