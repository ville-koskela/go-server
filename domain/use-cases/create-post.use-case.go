package usecases

import (
	"errors"

	"web1/domain/models"
)

type CreatePostUseCase struct {
	db Database
}

func NewCreatePostUseCase(db Database) *CreatePostUseCase {
	return &CreatePostUseCase{db: db}
}

func (createPost *CreatePostUseCase) Execute(post *models.Post) (models.Post, error) {
	newPost, err := createPost.db.SavePost(post)

	if err != nil {
		return models.Post{}, errors.New("Error saving post")
	}

	return newPost, nil

}
