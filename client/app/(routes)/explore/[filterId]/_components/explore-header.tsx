import { Button } from "@/components/ui/button";
import Link from "next/link";

type ExploreHeaderProps = {
  filterId: string;
  categoryId?: string;
};

const ExploreHeader: React.FC<ExploreHeaderProps> = ({
  filterId,
  categoryId,
}) => {
  const filters = ["videos", "channels", "podcast"];

  return (
    <div className="flex flex-col gap-2">
      <h1 className="text-2xl font-semibold">
        {categoryId ? categoryId.toUpperCase() : "SEMUANYA"}
      </h1>
      <div className="flex gap-4">
        {filters.map((item) => (
          <Button
            key={item}
            variant={item === filterId ? "default" : "ghost"}
            asChild
          >
            <Link href={`/explore/${item}/${categoryId ? categoryId : "all"}`}>
              <p>{item.toUpperCase()}</p>
            </Link>
          </Button>
        ))}
      </div>
    </div>
  );
};

export default ExploreHeader;
