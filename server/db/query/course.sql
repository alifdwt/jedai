-- name: CreateCourse :one
INSERT INTO courses (
  id,
  user_id,
  title,
  description,
  image_url,
  price,
  is_published,
  category_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetCourse :one
SELECT * FROM courses
INNER JOIN categories ON courses.category_id = categories.id
INNER JOIN users ON courses.user_id = users.username
WHERE courses.id = $1 AND courses.user_id = $2
ORDER BY courses.created_at DESC
LIMIT 1;

-- name: ListCourses :many
SELECT * FROM courses
INNER JOIN categories ON courses.category_id = categories.id
INNER JOIN users ON courses.user_id = users.username
ORDER BY courses.created_at DESC
LIMIT $1
OFFSET $2;

-- name: ListCoursesByUserID :many
SELECT * FROM courses
INNER JOIN categories ON courses.category_id = categories.id
INNER JOIN users ON courses.user_id = users.username
WHERE user_id = $1
ORDER BY courses.created_at DESC
LIMIT $2
OFFSET $3;

-- name: UpdateCourse :one
UPDATE courses
SET id = $3, title = $4, description = $5, image_url = $6, price = $7, is_published = $8, category_id = $9
WHERE id = $1
    AND user_id = $2
RETURNING *;

-- name: DeleteCourse :exec
DELETE FROM courses
WHERE id = $1;
