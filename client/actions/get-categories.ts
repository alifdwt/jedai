import { Category } from "@/types";

type GetCategories = {
  page_id: number;
  page_size: number;
};

export const getCategories = async ({
  page_id,
  page_size,
}: GetCategories): Promise<Category[]> => {
  const response = await fetch(
    `http://localhost:8080/categories?page_id=${page_id}&page_size=${page_size}`
  );
  return response.json();
};
