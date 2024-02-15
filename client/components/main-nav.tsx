"use client";

import { cn } from "@/lib/utils";
import { NavRoutes } from "@/types";
import Link from "next/link";
import { usePathname } from "next/navigation";
import React from "react";

export function MainNav({ className }: React.HTMLAttributes<HTMLElement>) {
  const pathname = usePathname();
  const routes: NavRoutes[] = [
    {
      href: "/",
      label: "Featured",
      active: pathname === "/",
    },
    {
      href: "/explore",
      label: "Explore",
      active: pathname.includes("/explore"),
    },
    {
      href: "/classes",
      label: "Classes",
      active: pathname.includes("/classes"),
    },
  ];

  return (
    <nav
      className={cn("mx-6 flex items-center space-x-4 lg:space-x-6", className)}
    >
      {routes.map((route) => (
        <Link
          key={route.href}
          href={route.href}
          className={cn(
            "text-sm font-medium transition-colors hover:text-white",
            route.active ? "text-white" : "text-neutral-500"
          )}
        >
          {route.label}
        </Link>
      ))}
    </nav>
  );
}
