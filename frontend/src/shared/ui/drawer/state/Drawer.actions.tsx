import { T_ActionMap } from "../../../types/action-map";

export enum E_DrawerAction {
  Toggle = "TOGGLE",
  Close = "CLOSE",
}

type T_DrawerActionPayload = {
  [E_DrawerAction.Toggle]: undefined;
  [E_DrawerAction.Close]: undefined;
};

export type T_DrawerAction =
  T_ActionMap<T_DrawerActionPayload>[keyof T_ActionMap<T_DrawerActionPayload>];

export enum E_DrawerState {
  Open = "OPEN",
  Close = "CLOSE",
}

export type T_DrawerState = {
  state: E_DrawerState;
};
