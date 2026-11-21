import { useContext } from "react";
import { DrawerContext } from "../state/Drawer.context";

export const useDrawer = () => {
  return useContext(DrawerContext);
};
