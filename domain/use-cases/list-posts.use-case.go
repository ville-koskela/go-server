package usecases

import (
	"errors"

	"web1/domain/interfaces"
	"web1/domain/models"
)

type ListPostsUseCase struct {
	db interfaces.IDatabase
}

func NewListPostsUseCase(db interfaces.IDatabase) *ListPostsUseCase {
	return &ListPostsUseCase{db: db}
}

func (listPosts *ListPostsUseCase) Execute() ([]models.Post, error) {
	posts, err := listPosts.db.ListPosts()

	if err != nil {
		return nil, errors.New("Error listing posts")
	}

	return posts, nil
}
