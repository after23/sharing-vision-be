package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/after23/sharing-vision-be/util"
	"github.com/stretchr/testify/require"
)

func createRandomPost(t *testing.T) (CreatePostParams,int64) {
	arg := CreatePostParams{
		Title: util.RandomTitle(),
		Content: util.RandomContent(),
		Category: util.RandomCategory(),
		Status: util.RandomStatus(),
		CreatedDate: sql.NullTime{
			Time: time.Now().UTC().Add(7*time.Hour),
			Valid: true,
		},
	}

	res, err := testQueries.CreatePost(context.Background(), arg)
	require.NoError(t, err)
	id, err := res.LastInsertId()
	require.NoError(t, err)
	require.NotEmpty(t, id)
	require.NotZero(t, id)

	return arg, id
}

func deleteRandomPost(t *testing.T, id int64) {
	res, err := testQueries.DeletePost(context.Background(), int32(id))
	require.NoError(t, err)
	rowsAffected, err := res.RowsAffected()
	require.NoError(t, err)
	require.NotZero(t, rowsAffected)
}

func TestCreatePost(t *testing.T) {
	postParam, id := createRandomPost(t)
	require.NotEmpty(t, postParam)
	require.NotZero(t, id)
	deleteRandomPost(t, id)
}

func TestGetPostById(t *testing.T) {
	postParam, id := createRandomPost(t)
	require.NotEmpty(t, postParam)
	require.NotZero(t, id)

	post, err := testQueries.GetPostById(context.Background(), int32(id))
	require.NoError(t, err)
	require.NotEmpty(t, post)

	require.Equal(t, post.Title, postParam.Title)
	require.Equal(t, post.Content, postParam.Content)
	require.Equal(t, post.Category, postParam.Category)
	require.Equal(t, post.Status, postParam.Status)
}

func TestGetPost(t *testing.T) {
	var ids []int64
	for i := 0 ; i < 10; i++ {
		_, id := createRandomPost(t)
		ids = append(ids, id)
	}

	arg := GetPostParams{
		Limit: 5,
		Offset: 1,
	}

	posts, err := testQueries.GetPost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, posts)
	require.Equal(t, len(posts), 5)

	for _,v := range posts {
		require.NotEmpty(t, v)
	}

	for _,v := range ids {
		deleteRandomPost(t,v)
	}
}

func TestUpdatePost(t *testing.T) {
	_, id := createRandomPost(t)
	require.NotZero(t, id)

	arg := UpdatePostParams{
		Title: util.RandomTitle(),
		Content: util.RandomCategory(),
		Status: util.RandomStatus(),
		Category: util.RandomCategory(),
		UpdatedDate: sql.NullTime{
			Time: time.Now().UTC().Add(7*time.Hour),
			Valid: true,
		},
		ID: int32(id),
	}

	res, err := testQueries.UpdatePost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	rowsAffected, err := res.RowsAffected()
	require.NoError(t, err)
	require.NotZero(t, rowsAffected)

	deleteRandomPost(t, id)
}

