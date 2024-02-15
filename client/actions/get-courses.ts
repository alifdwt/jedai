import { Course } from "@/types";

type GetCourses = {
  page_id: number;
  page_size: number;
  category_id?: string | null;
};

export const getCourses = async ({
  page_id,
  page_size,
  category_id,
}: GetCourses): Promise<Course[]> => {
  const response = await fetch(
    `http://localhost:8080/courses?page_id=${page_id}&page_size=${page_size}&category_id=${category_id}`
  );
  return response.json();
};
