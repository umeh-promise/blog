package services

import (
	"context"

	"github.com/umeh-promise/blog/internal/interfaces"
	"github.com/umeh-promise/blog/internal/models"
)

type PostService struct {
	Repo interfaces.Posts
}

func NewPostService(repo interfaces.Posts) *PostService {
	return &PostService{
		Repo: repo,
	}
}

func (service *PostService) Create(ctx context.Context, post *models.Post) error {
	return service.Repo.Create(ctx, post)
}

func (service *PostService) GetByID(ctx context.Context, id int64) (*models.Post, error) {
	return service.Repo.GetByID(ctx, id)
}

func (service *PostService) Update(ctx context.Context, post *models.Post) error {
	return service.Repo.Update(ctx, post)
}

func (service *PostService) Delete(ctx context.Context, id int64) error {
	return service.Repo.Delete(ctx, id)
}

func (service *PostService) GetAll(ctx context.Context) ([]models.Post, error) {
	return service.Repo.GetAll(ctx)
}
