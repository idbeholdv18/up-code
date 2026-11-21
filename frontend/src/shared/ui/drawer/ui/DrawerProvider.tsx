import { FC, PropsWithChildren, useReducer } from "react";
import { DrawerContext, DrawerDispatchContext } from "../state/Drawer.context";
import { drawerReducer } from "../state/Drawer.reducer";
import { E_DrawerState } from "../state/Drawer.actions";

type DrawerProps = PropsWithChildren<{
  defaultState?: E_DrawerState;
}>;

export const DrawerProvider: FC<DrawerProps> = (props) => {
  const [state, dispatch] = useReducer(drawerReducer, {
    state: props.defaultState || E_DrawerState.Close,
  });
  return (
    <DrawerContext.Provider value={state}>
      <DrawerDispatchContext.Provider value={dispatch}>
        {props.children}
      </DrawerDispatchContext.Provider>
    </DrawerContext.Provider>
  );
};
