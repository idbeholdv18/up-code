import clsx from "clsx";
import { FC, PropsWithChildren } from "react";

export type NavbarLayoutProps = PropsWithChildren<{
  className?: string;
}>;

export const NavbarLayout: FC<NavbarLayoutProps> = (props) => {
  return (
    <div
      className={clsx(
        "mx-auto h-full px-4 sm:px-6 md:px-8 flex items-center w-full",
        props.className,
      )}
    >
      {props.children}
    </div>
  );
};
