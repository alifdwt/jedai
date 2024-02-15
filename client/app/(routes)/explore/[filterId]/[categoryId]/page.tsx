import ExploreHeader from "../_components/explore-header";
import ExploreSection from "../_components/explore-section";

interface CategoryPageProps {
  params: {
    filterId: string;
    categoryId: string;
  };
}

const CategoryPage: React.FC<CategoryPageProps> = ({ params }) => {
  return (
    <div className="h-full p-8 flex flex-col gap-6 overflow-y-auto">
      <ExploreHeader
        filterId={params.filterId}
        categoryId={params.categoryId}
      />
      <ExploreSection
        filterId={params.filterId}
        categoryId={params.categoryId}
      />
    </div>
  );
};

export default CategoryPage;
