package interfaces

import (
	"context"

	"github.com/umeh-promise/blog/internal/models"
)

type Users interface {
	Create(context.Context, *models.User) error
	GetByID(context.Context, int64) (*models.User, error)
	GetByEmail(context.Context, string) (*models.User, error)
	Update(context.Context, *models.User) error
	Delete(context.Context, int64) error
}
