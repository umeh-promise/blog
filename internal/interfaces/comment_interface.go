package interfaces

import (
	"context"

	"github.com/umeh-promise/blog/internal/models"
)

type Comment interface {
	Create(context.Context, *models.Comment) error
	GetByPostID(context.Context, int64) ([]models.Comment, error)
}
