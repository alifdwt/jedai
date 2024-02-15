import ExploreChannel from "./explore-channel";
import ExploreCourses from "./explore-courses";

type ExploreSectionProps = {
  filterId: string;
  categoryId: string;
};

const ExploreSection: React.FC<ExploreSectionProps> = ({
  filterId,
  categoryId,
}) => {
  if (filterId === "videos") {
    return <ExploreCourses categoryId={categoryId} />;
  } else if (filterId === "channels") {
    return <ExploreChannel filterId={filterId} categoryId={categoryId} />;
  }
};

export default ExploreSection;
