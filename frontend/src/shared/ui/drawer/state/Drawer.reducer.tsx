import {
  E_DrawerAction,
  E_DrawerState,
  T_DrawerAction,
  T_DrawerState,
} from "./Drawer.actions";

export function drawerReducer(
  state: T_DrawerState,
  action: T_DrawerAction,
): T_DrawerState {
  switch (action.type) {
    case E_DrawerAction.Toggle: {
      return {
        ...state,
        state:
          state.state === E_DrawerState.Open
            ? E_DrawerState.Close
            : E_DrawerState.Open,
      };
    }
    case E_DrawerAction.Close: {
      return {
        ...state,
        state: E_DrawerState.Close,
      };
    }
    default:
      return state;
  }
}
