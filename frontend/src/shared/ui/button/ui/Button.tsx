import { type ButtonHTMLAttributes } from "react";
import clsx from "clsx";
import { Slot } from "../../slot";

interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  asChild?: boolean;
  variant?: "outline" | "solid";
}

export const Button = ({
  asChild = false,
  variant = "solid",
  className,
  ...props
}: ButtonProps) => {
  const Component = asChild ? Slot : "button";

  return (
    <Component
      data-variant={variant}
      className={clsx(
        "px-4 h-9 inline-flex items-center justify-center cursor-pointer",
        className,
      )}
      {...props}
    />
  );
};
