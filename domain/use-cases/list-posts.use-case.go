package usecases

import (
	"errors"

	"web1/domain/models"
)

func (uc *UseCases) ListPosts() ([]models.Post, error) {
	posts, err := uc.db.ListPosts()

	if err != nil {
		return nil, errors.New("Error listing posts")
	}

	return posts, nil
}
