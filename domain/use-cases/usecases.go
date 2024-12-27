package usecases

import (
	"web1/domain/models"
)

type Database interface {
	SavePost(post *models.Post) (models.Post, error)
	ListPosts() ([]models.Post, error)
	GetPost(id int64) (models.Post, error)
	SaveComment(comment *models.Comment) (models.Comment, error)
	ListComments(postID int64) ([]models.Comment, error)
}

type UseCases struct {
	db Database
}

func NewUseCases(db Database) *UseCases {
	return &UseCases{db: db}
}
