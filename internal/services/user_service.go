package services

import (
	"context"

	"github.com/umeh-promise/blog/internal/interfaces"
	"github.com/umeh-promise/blog/internal/models"
)

type UserService struct {
	Repo interfaces.Users
}

func NewUserService(repo interfaces.Users) *UserService {
	return &UserService{Repo: repo}
}

func (service *UserService) Create(ctx context.Context, user *models.User) error {
	return service.Repo.Create(ctx, user)
}

func (service *UserService) GetUserByID(ctx context.Context, id int64) (*models.User, error) {
	return service.Repo.GetByID(ctx, id)
}

func (service *UserService) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	return service.Repo.GetByEmail(ctx, email)
}

func (service *UserService) Update(ctx context.Context, user *models.User) error {
	return service.Repo.Update(ctx, user)
}

func (service *UserService) Delete(ctx context.Context, id int64) error {
	return service.Repo.Delete(ctx, id)
}
