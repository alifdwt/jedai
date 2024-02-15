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

// func TestGetUser(t *testing.T) {
// 	user1 := createRandomUser(t)
// 	user2, err := testQueries.GetUserWithCourses(context.Background(), user1.Username)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, user2)

// 	require.Equal(t, user1.Username, user2.Username)
// 	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
// 	require.Equal(t, user1.FullName, user2.FullName)
// 	require.Equal(t, user1.Email, user2.Email)

// 	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
// 	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
// }
