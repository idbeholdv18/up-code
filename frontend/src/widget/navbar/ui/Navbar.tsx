import { Button } from "@/shared/ui/button";
import {
  DrawerLayout,
  E_DrawerAction,
  useDispatchDrawer,
} from "@/shared/ui/drawer";
import { Navbar as SharedNavbar } from "@/shared/ui/navbar";
import { FC } from "react";
import { Link } from "react-router";
import { NavbarLink } from "./Navbar.Link";

export type NavbarProps = {};

export const Navbar: FC<NavbarProps> = (props) => {
  const dispatchDrawer = useDispatchDrawer();

  const toggleDrawer = () => {
    dispatchDrawer({ type: E_DrawerAction.Toggle });
  };

  const closeDrawer = () => {
    dispatchDrawer({ type: E_DrawerAction.Close });
  };

  return (
    <SharedNavbar.Base className='bg-bg-accent h-14 border-b border-b-neutral-800'>
      <SharedNavbar.Layout className='justify-between'>
        <SharedNavbar.Container className='flex gap-4 w-full'>
          {/* Logo */}
          <SharedNavbar.Container>
            <h1>P.Name</h1>
          </SharedNavbar.Container>

          {/* Panel */}
          <SharedNavbar.Container className='hidden md:flex gap-4 w-full justify-between'>
            {/* Navigation */}
            <SharedNavbar.Container>
              <NavbarLink to={"/"}>Home</NavbarLink>
            </SharedNavbar.Container>

            {/* Tools */}
            <SharedNavbar.Container className='hidden md:flex gap-4 self-end'>
              <Button
                asChild
                onClick={closeDrawer}
                className='text-fg-accent px-4 h-9 rounded-sm border border-fg-accent'
              >
                <Link to='/auth/login'>Sign In</Link>
              </Button>
              <Button
                asChild
                onClick={closeDrawer}
                className='text-bg-accent bg-fg-accent px-4 h-9 rounded-sm'
              >
                <Link to='/auth/register'>Sign Up</Link>
              </Button>
            </SharedNavbar.Container>
          </SharedNavbar.Container>
        </SharedNavbar.Container>
        <SharedNavbar.Container>
          <button onClick={toggleDrawer} className='text-fg-accent md:hidden'>
            Tog
          </button>
        </SharedNavbar.Container>
      </SharedNavbar.Layout>
      <DrawerLayout className='bg-bg-accent fixed top-0 w-screen z-0 px-4 md:hidden'>
        <SharedNavbar.Container className='min-h-screen pt-16 pb-8 flex flex-col justify-between'>
          <SharedNavbar.Container className='h-full w-full'>
            <NavbarLink to={"/"} onClick={closeDrawer}>
              Home
            </NavbarLink>
          </SharedNavbar.Container>

          <SharedNavbar.Container className='h-full flex flex-col w-full gap-4'>
            <Button
              asChild
              onClick={closeDrawer}
              className='text-fg-accent px-4 h-9 rounded-sm border border-fg-accent w-full'
            >
              <Link to='/auth/login'>Sign In</Link>
            </Button>
            <Button
              asChild
              onClick={closeDrawer}
              className='text-bg-accent bg-fg-accent px-4 h-9 rounded-sm w-full'
            >
              <Link to='/auth/register'>Sign Up</Link>
            </Button>
          </SharedNavbar.Container>
        </SharedNavbar.Container>
      </DrawerLayout>
    </SharedNavbar.Base>
  );
};
