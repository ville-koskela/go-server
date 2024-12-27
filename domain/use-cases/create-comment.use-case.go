package usecases

import (
	"errors"

	"web1/domain/interfaces"
	"web1/domain/models"
)

type CreateCommentUseCase struct {
	db interfaces.IDatabase
}

func NewCreateCommentUseCase(db interfaces.IDatabase) *CreateCommentUseCase {
	return &CreateCommentUseCase{db: db}
}

func (createComment *CreateCommentUseCase) Execute(comment *models.Comment) (models.Comment, error) {
	_, err := createComment.db.GetPost(comment.PostID)

	if err != nil {
		return models.Comment{}, errors.New("Post not found")
	}

	newComment, err := createComment.db.SaveComment(comment)

	if err != nil {
		return models.Comment{}, errors.New("Error saving comment")
	}

	return newComment, nil

}
