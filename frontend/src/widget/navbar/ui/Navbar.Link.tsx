import { FC } from "react";
import { Link, LinkProps } from "react-router";
import { clsx } from "clsx";

export const NavbarLink: FC<LinkProps> = ({ className, ...props }) => {
  return (
    <Link
      className={clsx("text-fg-accent h-9 flex items-center", className)}
      {...props}
    >
      {props.children}
    </Link>
  );
};
