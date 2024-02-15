-- name: CreateUser :one
INSERT INTO users (
    username,
    hashed_password,
    full_name,
    email,
    image_url,
    banner_url
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: ListUsers :many
SELECT * FROM users
LIMIT $1
OFFSET $2;

-- name: GetUserWithCourses :one
SELECT * FROM users_with_courses
WHERE username = $1;