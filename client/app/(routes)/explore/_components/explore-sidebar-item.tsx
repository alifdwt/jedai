"use client";

import { Button } from "@/components/ui/button";
import { Category } from "@/types";
import Link from "next/link";
import { usePathname } from "next/navigation";

type ExploreSidebarItemProps = {
  category: Category;
};

const ExploreSidebarItem: React.FC<ExploreSidebarItemProps> = ({
  category,
}) => {
  const pathname = usePathname();
  const isSelected = pathname === `/explore/videos/${category.id}`;

  return (
    <Button
      asChild
      variant={isSelected ? "secondary" : "ghost"}
      className="w-full"
    >
      <Link href={`/explore/videos/${category.id}`}>
        <p className="w-full text-left">{category.name}</p>
      </Link>
    </Button>
  );
};

export default ExploreSidebarItem;
