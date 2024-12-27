package database

import (
	"errors"
	"sync"

	"web1/domain/models"
)

type InMemoryDatabase struct {
	posts    map[int]models.Post
	comments map[int][]models.Comment
	postID   int
	commentID int
	mu       sync.Mutex
}

func NewInMemoryDatabase() *InMemoryDatabase {
	return &InMemoryDatabase{
		posts:    make(map[int]models.Post),
		comments: make(map[int][]models.Comment),
		postID:   0,
		commentID: 0,
	}
}

func (db *InMemoryDatabase) SavePost(post *models.Post) (models.Post, error) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.postID++
	post.ID = db.postID
	db.posts[db.postID] = *post
	return *post, nil
}

func (db *InMemoryDatabase) ListPosts() ([]models.Post, error) {
	db.mu.Lock()
	defer db.mu.Unlock()
	var posts []models.Post
	for _, post := range db.posts {
		posts = append(posts, post)
	}
	return posts, nil
}

func (db *InMemoryDatabase) GetPost(id int) (models.Post, error) {
	db.mu.Lock()
	defer db.mu.Unlock()
	post, exists := db.posts[id]
	if !exists {
		return models.Post{}, errors.New("post not found")
	}
	return post, nil
}

func (db *InMemoryDatabase) SaveComment(comment *models.Comment) (models.Comment, error) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.commentID++
	comment.ID = db.commentID
	db.comments[comment.PostID] = append(db.comments[comment.PostID], *comment)
	return *comment, nil
}

func (db *InMemoryDatabase) ListComments(postID int) ([]models.Comment, error) {
	db.mu.Lock()
	defer db.mu.Unlock()
	comments, exists := db.comments[postID]
	if !exists {
		return nil, errors.New("no comments for this post")
	}
	return comments, nil
}
