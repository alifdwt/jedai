package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/alifdwt/jedai/server/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       util.RandomString(6),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
		ImageUrl:       fmt.Sprintf("https://avatar.iran.liara.run/public/%d", util.RandomInt(1, 30)),
		BannerUrl:      fmt.Sprintf("https://picsum.photos/id/%d/200/300", util.RandomInt(1, 30)),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUserWithCourses(t *testing.T) {
	category := createRandomCategory(t)
	course := createRandomCourse(t, category.ID)
	user2, err := testQueries.GetUserWithCourses(context.Background(), course.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	arg := ListUsersParams{
		Limit:  5,
		Offset: 0,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, users)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}
