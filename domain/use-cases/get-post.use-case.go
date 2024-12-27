package usecases

import (
	"errors"

	"web1/domain/interfaces"
	"web1/domain/models"
)

type GetPostUseCase struct  {
	db interfaces.IDatabase
}

func NewGetPostUseCase(db interfaces.IDatabase) *GetPostUseCase {
	return &GetPostUseCase{db: db}
}

func (getPost *GetPostUseCase) Execute(id int64) (models.Post, []models.Comment, error) {
	post, err := getPost.db.GetPost(id)

	if err != nil {
		return models.Post{}, nil, errors.New("Post not found")
	}

	comments, err := getPost.db.ListComments(id)

	if err != nil {
		return models.Post{}, nil, errors.New("Error listing comments")
	}

	return post, comments, nil
}
