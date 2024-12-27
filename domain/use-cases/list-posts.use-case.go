package usecases

import (
	"errors"

	"web1/domain/models"
)

type ListPostsUseCase struct {
	db Database
}

func NewListPostsUseCase(db Database) *ListPostsUseCase {
	return &ListPostsUseCase{db: db}
}

func (listPosts *ListPostsUseCase) Execute() ([]models.Post, error) {
	posts, err := listPosts.db.ListPosts()

	if err != nil {
		return nil, errors.New("Error listing posts")
	}

	return posts, nil
}
