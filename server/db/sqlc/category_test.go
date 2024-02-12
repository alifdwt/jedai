package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/alifdwt/jedai/server/util"
	"github.com/stretchr/testify/require"
)

func createRandomCatogory(t *testing.T) Category {
	arg := CreateCategoryParams{
		ID:   util.RandomString(6),
		Name: util.RandomOwner(),
	}

	category, err := testQueries.CreateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, arg.ID, category.ID)
	require.Equal(t, arg.Name, category.Name)

	return category
}

func TestCreateCategory(t *testing.T) {
	createRandomCatogory(t)
}

func TestGetCategory(t *testing.T) {
	category1 := createRandomCatogory(t)
	category2, err := testQueries.GetCategory(context.Background(), category1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, category2.Name, category2.Name)
}

func TestUpdateCategory(t *testing.T) {
	category1 := createRandomCatogory(t)

	arg := UpdateCategoryParams{
		ID:   category1.ID,
		ID_2: util.RandomString(6),
		Name: util.RandomOwner(),
	}

	category2, err := testQueries.UpdateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.NotEqual(t, category1.ID, category2.ID)
	require.Equal(t, arg.ID_2, category2.ID)
	require.Equal(t, arg.Name, category2.Name)
}

func TestDeleteCategory(t *testing.T) {
	category1 := createRandomCatogory(t)
	err := testQueries.DeleteCategory(context.Background(), category1.ID)
	require.NoError(t, err)

	category2, err := testQueries.GetCategory(context.Background(), category1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, category2)
}

func TestListCategories(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCatogory(t)
	}

	arg := ListCategoriesParams{
		Limit:  5,
		Offset: 0,
	}

	categories, err := testQueries.ListCategories(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, categories)

	for _, category := range categories {
		require.NotEmpty(t, category)
	}
}