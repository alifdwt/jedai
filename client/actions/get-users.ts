import { User } from "@/types";

type GetUsers = {
  page_id: number;
  page_size: number;
  category_id?: string | null;
};

export const getUsers = async ({
  page_id,
  page_size,
  category_id,
}: GetUsers): Promise<User[]> => {
  const response = await fetch(
    `http://localhost:8080/users?page_id=${page_id}&page_size=${page_size}`
  );
  return response.json();
};
