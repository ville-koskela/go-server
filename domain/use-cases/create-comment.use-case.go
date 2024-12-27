package usecases

import (
	"errors"

	"web1/domain/models"
)

func (uc *UseCases) CreateComment(comment *models.Comment) (models.Comment, error) {
	_, err := uc.db.GetPost(comment.PostID)

	if err != nil {
		return models.Comment{}, errors.New("Post not found")
	}

	newComment, err := uc.db.SaveComment(comment)

	if err != nil {
		return models.Comment{}, errors.New("Error saving comment")
	}

	return newComment, nil
}
