"use client";

import { ShoppingBag } from "lucide-react";
import { useEffect, useState } from "react";
import { Button } from "./ui/button";

const NavbarAction = () => {
  const [isMounted, setMounted] = useState(false);

  useEffect(() => {
    setMounted(true);
  }, []);

  if (!isMounted) {
    return null;
  }
  return (
    <div className="ml-auto items-center gap-x-4">
      <Button
        variant={"secondary"}
        className="flex items-center rounded-full px-4 py-2"
      >
        <ShoppingBag size={20} />
        <span className="ml-2 text-sm font-medium">0</span>
      </Button>
    </div>
  );
};

export default NavbarAction;
