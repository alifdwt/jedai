export interface NavRoutes {
  href: string;
  label: string;
  active: boolean;
}

export interface Category {
  id: string;
  name: string;
}

export interface User {
  username: string;
  full_name: string;
  email: string;
  image_url: string;
  banner_url: string;
  password_changed_at: string;
  created_at: string;
}

export interface Course {
  id: string;
  user: User;
  title: string;
  description: string;
  image_url: string;
  price: number;
  is_published: boolean;
  category: Category;
  created_at: string;
}
