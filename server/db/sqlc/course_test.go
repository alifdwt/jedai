package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/alifdwt/jedai/server/util"
	"github.com/stretchr/testify/require"
)

func createRandomCourse(t *testing.T, categoryId string) Course {
	user := createRandomUser(t)
	arg := CreateCourseParams{
		ID:          util.RandomString(6),
		UserID:      user.Username,
		Title:       util.RandomOwner(),
		Price:       util.RandomMoney(),
		Description: util.RandomString(15),
		ImageUrl:    fmt.Sprintf("https://picsum.photos/id/%d/200/300", util.RandomInt(1, 30)),
		IsPublished: true,
		CategoryID:  categoryId,
	}

	course, err := testQueries.CreateCourse(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, course)

	require.Equal(t, arg.ID, course.ID)
	require.Equal(t, arg.UserID, course.UserID)
	require.Equal(t, arg.Title, course.Title)
	require.Equal(t, arg.Description, course.Description)
	require.Equal(t, arg.ImageUrl, course.ImageUrl)
	require.Equal(t, arg.IsPublished, course.IsPublished)
	require.Equal(t, arg.CategoryID, course.CategoryID)

	require.NotZero(t, course.CreatedAt)

	return course
}

func TestCreateCourse(t *testing.T) {
	category := createRandomCategory(t)
	createRandomCourse(t, category.ID)
}

func TestGetCourse(t *testing.T) {
	category := createRandomCategory(t)
	course1 := createRandomCourse(t, category.ID)
	arg := GetCourseParams{
		ID:     course1.ID,
		UserID: course1.UserID,
	}

	course2, err := testQueries.GetCourse(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, course2)

	require.Equal(t, course1.ID, course2.ID)
	require.Equal(t, course1.UserID, course2.UserID)
	require.Equal(t, course1.Title, course2.Title)
	require.Equal(t, course1.Description, course2.Description)
	require.Equal(t, course1.ImageUrl, course2.ImageUrl)
	require.Equal(t, course1.IsPublished, course2.IsPublished)
	require.Equal(t, course1.CategoryID, course2.CategoryID)

	require.WithinDuration(t, course1.CreatedAt, course2.CreatedAt, time.Second)
}

func TestListCourses(t *testing.T) {
	var lastCourse Course
	category := createRandomCategory(t)
	for i := 0; i < 10; i++ {
		lastCourse = createRandomCourse(t, category.ID)
	}

	arg := ListCoursesParams{
		Limit:      5,
		Offset:     0,
		CategoryID: fmt.Sprintf("%%%s%%", lastCourse.CategoryID),
	}

	courses, err := testQueries.ListCourses(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, courses)

	for _, course := range courses {
		require.NotEmpty(t, course)
	}
}

func TestListCoursesByUserID(t *testing.T) {
	var lastCourse Course
	category := createRandomCategory(t)
	for i := 0; i < 10; i++ {
		lastCourse = createRandomCourse(t, category.ID)
	}

	arg := ListCoursesByUserIDParams{
		UserID: lastCourse.UserID,
		Limit:  5,
		Offset: 0,
	}

	courses, err := testQueries.ListCoursesByUserID(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, courses)

	for _, course := range courses {
		require.NotEmpty(t, course)
		require.Equal(t, lastCourse.UserID, course.UserID)
	}
}

func TestUpdateCourse(t *testing.T) {
	category := createRandomCategory(t)
	course1 := createRandomCourse(t, category.ID)

	arg := UpdateCourseParams{
		ID:          course1.ID,
		ID_2:        util.RandomString(6),
		UserID:      course1.UserID,
		Title:       util.RandomOwner(),
		Price:       util.RandomMoney(),
		Description: util.RandomString(15),
		ImageUrl:    fmt.Sprintf("https://placehold.co/600x400?text=%s", course1.UserID),
		IsPublished: true,
		CategoryID:  course1.CategoryID,
	}

	course2, err := testQueries.UpdateCourse(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, course2)

	require.NotEqual(t, course1.ID, course2.ID)
	require.Equal(t, arg.ID_2, course2.ID)
	require.Equal(t, arg.Title, course2.Title)
	require.Equal(t, arg.Description, course2.Description)
	require.Equal(t, arg.ImageUrl, course2.ImageUrl)
	require.Equal(t, arg.IsPublished, course2.IsPublished)
	require.Equal(t, arg.CategoryID, course2.CategoryID)

	require.NotZero(t, course2.UpdatedAt)
	require.WithinDuration(t, course1.CreatedAt, course2.CreatedAt, time.Second)
}

func TestDeleteCourse(t *testing.T) {
	category := createRandomCategory(t)
	course1 := createRandomCourse(t, category.ID)
	err := testQueries.DeleteCourse(context.Background(), course1.ID)
	require.NoError(t, err)

	arg := GetCourseParams{
		ID:     course1.ID,
		UserID: course1.UserID,
	}

	course2, err := testQueries.GetCourse(context.Background(), arg)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, course2)
}
