package repositories

import (
	"context"
	"database/sql"

	"github.com/umeh-promise/blog/internal/interfaces"
	"github.com/umeh-promise/blog/internal/models"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) interfaces.Posts {
	return &PostRepository{db: db}
}

func (postRepo *PostRepository) Create(ctx context.Context, post *models.Post) error {

	return nil
}

func (postRepo *PostRepository) GetByID(ctx context.Context, id int64) (*models.Post, error) {

	return &models.Post{}, nil
}
