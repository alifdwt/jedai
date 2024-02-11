CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "courses" (
  "id" varchar PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "user_id" varchar NOT NULL,
  "title" varchar NOT NULL,
  "description" varchar,
  "image_url" varchar,
  "price" bigint,
  "is_published" bool NOT NULL,
  "category_id" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "categories" (
  "id" varchar PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" varchar NOT NULL
);

CREATE TABLE "attachments" (
  "id" varchar PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" varchar NOT NULL,
  "url" varchar NOT NULL,
  "course_id" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "chapters" (
  "id" varchar PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "title" varchar NOT NULL,
  "description" varchar,
  "video_url" varchar,
  "position" bigint NOT NULL,
  "is_published" bool NOT NULL DEFAULT false,
  "is_free" bool NOT NULL DEFAULT false,
  "course_id" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "muxes" (
  "id" varchar PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "asset_id" varchar NOT NULL,
  "playback_id" varchar,
  "chapter_id" varchar UNIQUE NOT NULL
);

CREATE TABLE "user_progress" (
  "id" varchar PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "user_id" varchar NOT NULL,
  "chapter_id" varchar NOT NULL,
  "is_completed" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "purchases" (
  "id" varchar PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "user_id" varchar NOT NULL,
  "course_id" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "stripe_customers" (
  "id" varchar PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "user_id" varchar NOT NULL,
  "stripe_customer_id" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "courses" ("category_id");

CREATE INDEX ON "attachments" ("course_id");

CREATE INDEX ON "chapters" ("course_id");

CREATE INDEX ON "user_progress" ("chapter_id");

CREATE UNIQUE INDEX ON "user_progress" ("user_id", "chapter_id");

CREATE INDEX ON "purchases" ("course_id");

CREATE UNIQUE INDEX ON "purchases" ("user_id", "course_id");

ALTER TABLE "courses" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("username");

ALTER TABLE "courses" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "attachments" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("id");

ALTER TABLE "chapters" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("id");

ALTER TABLE "muxes" ADD FOREIGN KEY ("chapter_id") REFERENCES "chapters" ("id");

ALTER TABLE "user_progress" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("username");

ALTER TABLE "user_progress" ADD FOREIGN KEY ("chapter_id") REFERENCES "chapters" ("id");

ALTER TABLE "purchases" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("username");

ALTER TABLE "purchases" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("id");

ALTER TABLE "stripe_customers" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("username");