import { createContext, Dispatch, useContext } from "react";
import { E_DrawerState, T_DrawerAction, T_DrawerState } from "./Drawer.actions";

export const DrawerContext = createContext<T_DrawerState>({
  state: E_DrawerState.Close,
});

export const DrawerDispatchContext = createContext<Dispatch<T_DrawerAction>>(
  () => {},
);
