package interfaces

import (
	"context"

	"github.com/umeh-promise/blog/internal/models"
)

type Role interface {
	GetByName(context.Context, string) (*models.Role, error)
}
