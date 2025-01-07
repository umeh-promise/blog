package interfaces

import (
	"context"

	"github.com/umeh-promise/blog/internal/models"
)

type Posts interface {
	Create(context.Context, *models.Post) error
	GetByID(context.Context, int64) (*models.Post, error)
}
