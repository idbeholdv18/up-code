import { DrawerLayout, DrawerProvider } from "@/shared/ui/drawer";
import { useDispatchDrawer } from "@/shared/ui/drawer/hooks/useDispatchDrawer";
import {
  E_DrawerAction,
  E_DrawerState,
} from "@/shared/ui/drawer/state/Drawer.actions";
import { Navbar } from "@/widget/navbar";
import { FC, PropsWithChildren } from "react";
import { Link, Outlet } from "react-router";

type LayoutProps = {};

export const RootLayout: FC<LayoutProps> = (props) => {
  return (
    <div className='pt-14'>
      <DrawerProvider>
        <Navbar />
      </DrawerProvider>

      {/* Content */}
      <Outlet />
    </div>
  );
};
