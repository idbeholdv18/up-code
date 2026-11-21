import { useContext } from "react";
import { DrawerDispatchContext } from "../state/Drawer.context";

export const useDispatchDrawer = () => {
  return useContext(DrawerDispatchContext);
};
