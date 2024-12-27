package usecases

import (
	"errors"

	"web1/domain/models"
)

func (uc *UseCases) GetPost(id int64) (models.Post, []models.Comment, error) {
	post, err := uc.db.GetPost(id)

	if err != nil {
		return models.Post{}, nil, errors.New("Post not found")
	}

	comments, err := uc.db.ListComments(id)

	if err != nil {
		return models.Post{}, nil, errors.New("Error listing comments")
	}

	return post, comments, nil
}
