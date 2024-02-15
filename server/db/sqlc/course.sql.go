// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: course.sql

package db

import (
	"context"
	"time"
)

const createCourse = `-- name: CreateCourse :one
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
) RETURNING id, user_id, title, description, image_url, price, is_published, category_id, created_at, updated_at
`

type CreateCourseParams struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
	Price       int64  `json:"price"`
	IsPublished bool   `json:"is_published"`
	CategoryID  string `json:"category_id"`
}

func (q *Queries) CreateCourse(ctx context.Context, arg CreateCourseParams) (Course, error) {
	row := q.db.QueryRowContext(ctx, createCourse,
		arg.ID,
		arg.UserID,
		arg.Title,
		arg.Description,
		arg.ImageUrl,
		arg.Price,
		arg.IsPublished,
		arg.CategoryID,
	)
	var i Course
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Description,
		&i.ImageUrl,
		&i.Price,
		&i.IsPublished,
		&i.CategoryID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteCourse = `-- name: DeleteCourse :exec
DELETE FROM courses
WHERE id = $1
`

func (q *Queries) DeleteCourse(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteCourse, id)
	return err
}

const getCourse = `-- name: GetCourse :one
SELECT courses.id, user_id, title, description, courses.image_url, price, is_published, category_id, courses.created_at, updated_at, categories.id, name, username, hashed_password, full_name, email, users.image_url, banner_url, password_changed_at, users.created_at FROM courses
INNER JOIN categories ON courses.category_id = categories.id
INNER JOIN users ON courses.user_id = users.username
WHERE courses.id = $1 AND courses.user_id = $2
ORDER BY courses.created_at DESC
LIMIT 1
`

type GetCourseParams struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
}

type GetCourseRow struct {
	ID                string    `json:"id"`
	UserID            string    `json:"user_id"`
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	ImageUrl          string    `json:"image_url"`
	Price             int64     `json:"price"`
	IsPublished       bool      `json:"is_published"`
	CategoryID        string    `json:"category_id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	ID_2              string    `json:"id_2"`
	Name              string    `json:"name"`
	Username          string    `json:"username"`
	HashedPassword    string    `json:"hashed_password"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	ImageUrl_2        string    `json:"image_url_2"`
	BannerUrl         string    `json:"banner_url"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt_2       time.Time `json:"created_at_2"`
}

func (q *Queries) GetCourse(ctx context.Context, arg GetCourseParams) (GetCourseRow, error) {
	row := q.db.QueryRowContext(ctx, getCourse, arg.ID, arg.UserID)
	var i GetCourseRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Description,
		&i.ImageUrl,
		&i.Price,
		&i.IsPublished,
		&i.CategoryID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ID_2,
		&i.Name,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.ImageUrl_2,
		&i.BannerUrl,
		&i.PasswordChangedAt,
		&i.CreatedAt_2,
	)
	return i, err
}

const listCourses = `-- name: ListCourses :many
SELECT courses.id, user_id, title, description, courses.image_url, price, is_published, category_id, courses.created_at, updated_at, categories.id, name, username, hashed_password, full_name, email, users.image_url, banner_url, password_changed_at, users.created_at FROM courses
INNER JOIN categories ON courses.category_id = categories.id
INNER JOIN users ON courses.user_id = users.username
WHERE courses.category_id LIKE $3
ORDER BY courses.created_at DESC
LIMIT $1
OFFSET $2
`

type ListCoursesParams struct {
	Limit      int32  `json:"limit"`
	Offset     int32  `json:"offset"`
	CategoryID string `json:"category_id"`
}

type ListCoursesRow struct {
	ID                string    `json:"id"`
	UserID            string    `json:"user_id"`
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	ImageUrl          string    `json:"image_url"`
	Price             int64     `json:"price"`
	IsPublished       bool      `json:"is_published"`
	CategoryID        string    `json:"category_id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	ID_2              string    `json:"id_2"`
	Name              string    `json:"name"`
	Username          string    `json:"username"`
	HashedPassword    string    `json:"hashed_password"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	ImageUrl_2        string    `json:"image_url_2"`
	BannerUrl         string    `json:"banner_url"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt_2       time.Time `json:"created_at_2"`
}

func (q *Queries) ListCourses(ctx context.Context, arg ListCoursesParams) ([]ListCoursesRow, error) {
	rows, err := q.db.QueryContext(ctx, listCourses, arg.Limit, arg.Offset, arg.CategoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListCoursesRow{}
	for rows.Next() {
		var i ListCoursesRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Description,
			&i.ImageUrl,
			&i.Price,
			&i.IsPublished,
			&i.CategoryID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ID_2,
			&i.Name,
			&i.Username,
			&i.HashedPassword,
			&i.FullName,
			&i.Email,
			&i.ImageUrl_2,
			&i.BannerUrl,
			&i.PasswordChangedAt,
			&i.CreatedAt_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCoursesByUserID = `-- name: ListCoursesByUserID :many
SELECT courses.id, user_id, title, description, courses.image_url, price, is_published, category_id, courses.created_at, updated_at, categories.id, name, username, hashed_password, full_name, email, users.image_url, banner_url, password_changed_at, users.created_at FROM courses
INNER JOIN categories ON courses.category_id = categories.id
INNER JOIN users ON courses.user_id = users.username
WHERE user_id = $1
ORDER BY courses.created_at DESC
LIMIT $2
OFFSET $3
`

type ListCoursesByUserIDParams struct {
	UserID string `json:"user_id"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

type ListCoursesByUserIDRow struct {
	ID                string    `json:"id"`
	UserID            string    `json:"user_id"`
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	ImageUrl          string    `json:"image_url"`
	Price             int64     `json:"price"`
	IsPublished       bool      `json:"is_published"`
	CategoryID        string    `json:"category_id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	ID_2              string    `json:"id_2"`
	Name              string    `json:"name"`
	Username          string    `json:"username"`
	HashedPassword    string    `json:"hashed_password"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	ImageUrl_2        string    `json:"image_url_2"`
	BannerUrl         string    `json:"banner_url"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt_2       time.Time `json:"created_at_2"`
}

func (q *Queries) ListCoursesByUserID(ctx context.Context, arg ListCoursesByUserIDParams) ([]ListCoursesByUserIDRow, error) {
	rows, err := q.db.QueryContext(ctx, listCoursesByUserID, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListCoursesByUserIDRow{}
	for rows.Next() {
		var i ListCoursesByUserIDRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Description,
			&i.ImageUrl,
			&i.Price,
			&i.IsPublished,
			&i.CategoryID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ID_2,
			&i.Name,
			&i.Username,
			&i.HashedPassword,
			&i.FullName,
			&i.Email,
			&i.ImageUrl_2,
			&i.BannerUrl,
			&i.PasswordChangedAt,
			&i.CreatedAt_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCourse = `-- name: UpdateCourse :one
UPDATE courses
SET id = $3, title = $4, description = $5, image_url = $6, price = $7, is_published = $8, category_id = $9
WHERE id = $1
    AND user_id = $2
RETURNING id, user_id, title, description, image_url, price, is_published, category_id, created_at, updated_at
`

type UpdateCourseParams struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	ID_2        string `json:"id_2"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
	Price       int64  `json:"price"`
	IsPublished bool   `json:"is_published"`
	CategoryID  string `json:"category_id"`
}

func (q *Queries) UpdateCourse(ctx context.Context, arg UpdateCourseParams) (Course, error) {
	row := q.db.QueryRowContext(ctx, updateCourse,
		arg.ID,
		arg.UserID,
		arg.ID_2,
		arg.Title,
		arg.Description,
		arg.ImageUrl,
		arg.Price,
		arg.IsPublished,
		arg.CategoryID,
	)
	var i Course
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Description,
		&i.ImageUrl,
		&i.Price,
		&i.IsPublished,
		&i.CategoryID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
