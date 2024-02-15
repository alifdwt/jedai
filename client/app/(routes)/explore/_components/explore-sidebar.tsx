import { Category } from "@/types";
import ExploreSidebarItem from "./explore-sidebar-item";

interface ExploreSidebarProps {
  categories: Category[];
}

const ExploreSidebar: React.FC<ExploreSidebarProps> = ({ categories }) => {
  return (
    <div className="h-full border-r flex flex-col overflow-y-auto shadow-sm">
      <div className="p-8 flex flex-col border-b">
        <h1 className="font-semibold">Kategori</h1>
      </div>
      <div className="flex flex-col w-full">
        {categories.map((category) => (
          <ExploreSidebarItem key={category.id} category={category} />
        ))}
      </div>
    </div>
  );
};

export default ExploreSidebar;
