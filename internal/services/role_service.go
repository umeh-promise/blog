package services

import (
	"context"

	"github.com/umeh-promise/blog/internal/interfaces"
	"github.com/umeh-promise/blog/internal/models"
)

type RoleService struct {
	Repo interfaces.Role
}

func NewRoleService(repo interfaces.Role) *RoleService {
	return &RoleService{Repo: repo}
}

func (service *RoleService) GetByName(ctx context.Context, name string) (*models.Role, error) {
	return service.Repo.GetByName(ctx, name)
}
