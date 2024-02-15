import { getCategories } from "@/actions/get-categories";
import ExploreSidebar from "./_components/explore-sidebar";

const ExploreLayout = async ({ children }: { children: React.ReactNode }) => {
  const categories = await getCategories({ page_id: 1, page_size: 30 });
  return (
    <div className="h-[93vh]">
      <div className="hidden md:flex h-full w-80 flex-col fixed z-50">
        <ExploreSidebar categories={categories} />
      </div>
      <main className="md:pl-80 h-full">{children}</main>
    </div>
  );
};

export default ExploreLayout;
