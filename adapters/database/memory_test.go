package database

import (
	"errors"
	"testing"

	"web1/domain/models"
	"web1/test"
)

func TestInMemoryDatabase_SavePost(t *testing.T) {
	db := NewInMemoryDatabase()

	post := &models.Post{Content: "Test Post"}
	savedPost, err := db.SavePost(post)

	test.Assert(t, nil, err)
	test.Assert(t, int64(1), savedPost.ID)
	test.Assert(t, "Test Post", savedPost.Content)
}

func TestInMemoryDatabase_ListPosts(t *testing.T) {
	db := NewInMemoryDatabase()

	post1 := &models.Post{Content: "Test Post 1"}
	post2 := &models.Post{Content: "Test Post 2"}
	db.SavePost(post1)
	db.SavePost(post2)

	posts, err := db.ListPosts()

	test.Assert(t, nil, err)
	test.Assert(t, 2, len(posts))
	test.Assert(t, "Test Post 1", posts[0].Content)
	test.Assert(t, "Test Post 2", posts[1].Content)
}

func TestInMemoryDatabase_GetPost(t *testing.T) {
	db := NewInMemoryDatabase()

	post := &models.Post{Content: "Test Post"}
	savedPost, _ := db.SavePost(post)

	retrievedPost, err := db.GetPost(savedPost.ID)

	test.Assert(t, nil, err)
	test.Assert(t, savedPost.ID, retrievedPost.ID)
	test.Assert(t, savedPost.Content, retrievedPost.Content)

	_, err = db.GetPost(999)
	test.Assert(t, errors.New("post not found"), err)
}

func TestInMemoryDatabase_SaveComment(t *testing.T) {
	db := NewInMemoryDatabase()

	post := &models.Post{Content: "Test Post"}
	savedPost, _ := db.SavePost(post)

	comment := &models.Comment{PostID: savedPost.ID, Content: "Test Comment"}
	savedComment, err := db.SaveComment(comment)

	test.Assert(t, nil, err)
	test.Assert(t, int64(1), savedComment.ID)
	test.Assert(t, savedPost.ID, savedComment.PostID)
	test.Assert(t, "Test Comment", savedComment.Content)
}

func TestInMemoryDatabase_ListComments(t *testing.T) {
	db := NewInMemoryDatabase()

	post := &models.Post{Content: "Test Post"}
	savedPost, _ := db.SavePost(post)

	comment1 := &models.Comment{PostID: savedPost.ID, Content: "Test Comment 1"}
	comment2 := &models.Comment{PostID: savedPost.ID, Content: "Test Comment 2"}
	db.SaveComment(comment1)
	db.SaveComment(comment2)

	comments, err := db.ListComments(savedPost.ID)

	test.Assert(t, nil, err)
	test.Assert(t, 2, len(comments))
	test.Assert(t, "Test Comment 1", comments[0].Content)
	test.Assert(t, "Test Comment 2", comments[1].Content)

	_, err = db.ListComments(999)
	test.Assert(t, errors.New("no comments for this post"), err)
}
