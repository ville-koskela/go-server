package database

import (
	"errors"
	"testing"

	"web1/domain/models"
	"web1/test"
)

var dbTypes = []string{"inmemory", "sqlite3"}

type MockEnv struct {
	dbType string
}

func (env MockEnv) GetDBType() string {
	return env.dbType
}

func (env MockEnv) GetDBPath() string {
	return ":memory:"
}

func TestDatabase_SavePost(t *testing.T) {
	for _, dbType := range dbTypes {
		t.Run(dbType, func(t *testing.T) {
			db := InitializeDatabase(MockEnv{dbType: dbType})
			defer db.Close()

			post := &models.Post{Content: "Test Post", Name: "Test Name"}
			savedPost, err := db.SavePost(post)

			test.Assert(t, nil, err)
			test.Assert(t, int64(1), savedPost.ID)
			test.Assert(t, "Test Post", savedPost.Content)
			test.Assert(t, "Test Name", savedPost.Name)
		})
	}
}

func TestDatabase_ListPosts(t *testing.T) {
	for _, dbType := range dbTypes {
		t.Run(dbType, func(t *testing.T) {
			db := InitializeDatabase(MockEnv{dbType: dbType})
			defer db.Close()

			post1 := &models.Post{Content: "Test Post 1"}
			post2 := &models.Post{Content: "Test Post 2"}
			db.SavePost(post1)
			db.SavePost(post2)

			posts, err := db.ListPosts()

			test.Assert(t, nil, err)
			test.Assert(t, 2, len(posts))
			test.Assert(t, "Test Post 2", posts[0].Content)
			test.Assert(t, "Test Post 1", posts[1].Content)
		})
	}
}

func TestDatabase_GetPost(t *testing.T) {
	for _, dbType := range dbTypes {
		t.Run(dbType, func(t *testing.T) {
			db := InitializeDatabase(MockEnv{dbType: dbType})
			defer db.Close()

			post := &models.Post{Content: "Test Post"}
			savedPost, _ := db.SavePost(post)

			retrievedPost, err := db.GetPost(savedPost.ID)

			test.Assert(t, nil, err)
			test.Assert(t, savedPost.ID, retrievedPost.ID)
			test.Assert(t, savedPost.Content, retrievedPost.Content)

			_, err = db.GetPost(999)
			test.Assert(t, errors.New("post not found"), err)
		})
	}
}

func TestDatabase_SaveComment(t *testing.T) {
	for _, dbType := range dbTypes {
		t.Run(dbType, func(t *testing.T) {
			db := InitializeDatabase(MockEnv{dbType: dbType})
			defer db.Close()

			post := &models.Post{Content: "Test Post"}
			savedPost, _ := db.SavePost(post)

			comment := &models.Comment{PostID: savedPost.ID, Content: "Test Comment"}
			savedComment, err := db.SaveComment(comment)

			test.Assert(t, nil, err)
			test.Assert(t, int64(1), savedComment.ID)
			test.Assert(t, savedPost.ID, savedComment.PostID)
			test.Assert(t, "Test Comment", savedComment.Content)
		})
	}
}

func TestInMemoryDatabase_ListComments(t *testing.T) {
	for _, dbType := range dbTypes {
		t.Run(dbType, func(t *testing.T) {
			db := InitializeDatabase(MockEnv{dbType: dbType})
			defer db.Close()

			post := &models.Post{Content: "Test Post"}
			savedPost, _ := db.SavePost(post)

			comment1 := &models.Comment{PostID: savedPost.ID, Content: "Test Comment 1", Name: "Test Name 1"}
			comment2 := &models.Comment{PostID: savedPost.ID, Content: "Test Comment 2", Name: "Test Name 2"}
			db.SaveComment(comment1)
			db.SaveComment(comment2)

			comments, err := db.ListComments(savedPost.ID)

			test.Assert(t, nil, err)
			test.Assert(t, 2, len(comments))
			test.Assert(t, "Test Comment 1", comments[0].Content)
			test.Assert(t, "Test Comment 2", comments[1].Content)
			test.Assert(t, "Test Name 1", comments[0].Name)
			test.Assert(t, "Test Name 2", comments[1].Name)

			_, err = db.ListComments(999)
			test.Assert(t, nil, err)
		})
	}
}
