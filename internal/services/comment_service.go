package services

import (
	"context"

	"github.com/umeh-promise/blog/internal/interfaces"
	"github.com/umeh-promise/blog/internal/models"
)

type CommentService struct {
	Repo interfaces.Comment
}

func NewCommentService(repo interfaces.Comment) *CommentService {
	return &CommentService{Repo: repo}
}

func (service *CommentService) Create(ctx context.Context, comment *models.Comment) error {
	return service.Repo.Create(ctx, comment)
}

func (service *CommentService) GetCommentByPostID(ctx context.Context, id int64) ([]models.Comment, error) {
	return service.Repo.GetByPostID(ctx, id)
}
