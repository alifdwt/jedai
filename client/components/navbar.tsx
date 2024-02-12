import Link from "next/link";
import Container from "./ui/container";
import { BrainCircuit } from "lucide-react";
import { MainNav } from "./main-nav";
import NavbarAction from "./navbar-action";

export const revalidate = 0;

const Navbar = () => {
  return (
    <div className="bg-primary">
      <Container>
        <div className="relative px-4 sm:px-6 lg:px-8 flex h-16 items-center">
          <Link href="/" className="ml-4 flex lg:ml-0 gap-x-2 items-center">
            <BrainCircuit size={24} className="text-secondary" />
            <p className="font-bold text-xl text-white">Jedai</p>
          </Link>
          <MainNav className="mx-4" />
          <NavbarAction />
        </div>
      </Container>
    </div>
  );
};

export default Navbar;
