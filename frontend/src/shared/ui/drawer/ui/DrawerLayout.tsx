import { createPortal } from "react-dom";
import { useDrawer } from "../hooks/useDrawer";
import { E_DrawerState } from "../state/Drawer.actions";
import { FC, PropsWithChildren } from "react";

type DrawerLayoutProps = PropsWithChildren<{
  className?: string;
}>;

export const DrawerLayout: FC<DrawerLayoutProps> = (props) => {
  const { state } = useDrawer();

  return (
    state === E_DrawerState.Open &&
    createPortal(<div className={props.className}>{props.children}</div>, document.body)
  );
};
