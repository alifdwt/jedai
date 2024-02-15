import ExploreCourses from "./_components/explore-courses";
import ExploreHeader from "./_components/explore-header";

type ExploreCategoryProps = {
  params: {
    filterId: string;
  };
};

const ExploreCategory: React.FC<ExploreCategoryProps> = ({ params }) => {
  return (
    <div className="h-full p-8 flex flex-col gap-6 overflow-y-auto">
      <ExploreHeader filterId={params.filterId} />
      <ExploreCourses filterId={params.filterId} categoryId="" />
    </div>
  );
};

export default ExploreCategory;
