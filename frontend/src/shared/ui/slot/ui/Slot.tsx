import {
  Children,
  cloneElement,
  isValidElement,
  type HTMLAttributes,
  type ReactNode,
  type Ref,
} from "react";
import { mergeProps } from "../lib/mergeProps";

function composeRefs<T>(...refs: (Ref<T> | undefined)[]) {
  return (node: T) => {
    refs.forEach((ref) => {
      if (typeof ref === "function") ref(node);
      else if (ref != null) (ref as React.MutableRefObject<T>).current = node;
    });
  };
}

interface SlotProps extends HTMLAttributes<HTMLElement> {
  children?: ReactNode;
  ref?: Ref<HTMLElement>;
}

export const Slot = ({ children, ref, ...slotProps }: SlotProps) => {
  if (!isValidElement(children)) {
    return null;
  }

  const child = Children.only(children) as React.ReactElement<any>;
  const childRef = child.props.ref ?? (child as any).ref;

  return cloneElement(child, {
    ...mergeProps(slotProps, child.props),
    ref: ref || childRef ? composeRefs(ref, childRef) : undefined,
  });
};
