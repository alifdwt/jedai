import { getCourses } from "@/actions/get-courses";
import { getDurationFromNow } from "@/lib/utils";
import Link from "next/link";

type ExploreCoursesProps = {
  categoryId?: string;
};

const ExploreCourses: React.FC<ExploreCoursesProps> = async ({
  categoryId,
}) => {
  const courses = await getCourses({
    page_id: 1,
    page_size: 12,
    category_id: categoryId === "all" ? "" : categoryId,
  });
  return (
    <section className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-x-4 gap-y-8">
      {courses ? (
        courses.map((item) => (
          <Link
            href={`/${item.user.username}/${item.id}`}
            className="w-full h-[250px]"
            key={item.id}
          >
            <div className="h-[200px]">
              <img
                src={item.image_url}
                alt={item.title}
                className="h-48 w-96 object-cover rounded-lg hover:scale-105 transition-all"
              />
            </div>
            <div className="h-[50px] flex items-center gap-2">
              <img
                src={item.user.image_url}
                alt={item.user.full_name}
                className="h-8 w-8 rounded-full"
              />
              <div>
                <p className="font-semibold text-sm">{item.title}</p>
                <p className="text-sm text-muted-foreground">
                  {item.user.full_name} •{" "}
                  {getDurationFromNow(new Date(item.created_at))}
                </p>
              </div>
            </div>
          </Link>
        ))
      ) : (
        <div className="w-full h-[250px] flex items-center justify-center">
          <h1 className="text-2xl font-semibold">Belum ada kelas</h1>
        </div>
      )}
    </section>
  );
};

export default ExploreCourses;
