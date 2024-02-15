import { getUsers } from "@/actions/get-users";
import Link from "next/link";

type ExploreChannelProps = {
  filterId: string;
  categoryId: string;
};

const ExploreChannel: React.FC<ExploreChannelProps> = async ({
  categoryId,
}) => {
  const channels = await getUsers({
    page_id: 1,
    page_size: 12,
    category_id: categoryId,
  });

  return (
    <section className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-x-4 gap-y-8">
      {channels.map((item) => (
        <Link
          href={`/${item.username}`}
          key={item.username}
          className="w-full h-[250px]"
        >
          <div className="h-[200px]">
            <img
              src={item.banner_url}
              alt={item.username}
              className="h-48 w-96 object-cover rounded-lg hover:scale-105 transition-all"
            />
          </div>
          <div className="h-[50px] flex items-center gap-2">
            <img
              src={item.image_url}
              alt={item.username}
              className="h-8 w-8 rounded-full"
            />
            <div>
              <p className="text-sm font-semibold">{item.username}</p>
              <p className="text-sm">{item.full_name}</p>
            </div>
          </div>
        </Link>
      ))}
    </section>
  );
};

export default ExploreChannel;
