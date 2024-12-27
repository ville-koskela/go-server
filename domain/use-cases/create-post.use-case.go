package usecases

import (
	"errors"

	"web1/domain/models"
)

func (uc *UseCases) CreatePost(post *models.Post) (models.Post, error) {
	newPost, err := uc.db.SavePost(post)

	if err != nil {
		return models.Post{}, errors.New("Error saving post")
	}

	return newPost, nil
}
