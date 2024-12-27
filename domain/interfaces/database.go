package interfaces

import (
	"web1/domain/models"
)

type IDatabase interface {
	SavePost(post *models.Post) (models.Post, error)
	ListPosts() ([]models.Post, error)
	GetPost(id int) (models.Post, error)
	SaveComment(comment *models.Comment) (models.Comment, error)
	ListComments(postID int) ([]models.Comment, error)
}
