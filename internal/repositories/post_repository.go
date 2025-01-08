package repositories

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
	"github.com/umeh-promise/blog/internal/interfaces"
	"github.com/umeh-promise/blog/internal/models"
	"github.com/umeh-promise/blog/internal/utils"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) interfaces.Posts {
	return &PostRepository{db: db}
}

func (postRepo *PostRepository) Create(ctx context.Context, post *models.Post) error {
	query := `INSERT INTO posts (user_id, title, content, tags) VALUES($1, $2, $3, $4) RETURNING id, created_at, updated_at`

	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	err := postRepo.db.QueryRowContext(ctx, query, post.UserID, post.Title, post.Content, pq.Array(post.Tags)).Scan(
		&post.ID,
		&post.CratedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (postRepo *PostRepository) GetByID(ctx context.Context, id int64) (*models.Post, error) {

	var post models.Post

	query := `SELECT id, user_id, title, content, tags, version, created_at, updated_at FROM posts
					WHERE id=$1`

	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	err := postRepo.db.QueryRowContext(ctx, query, id).Scan(
		&post.ID,
		&post.UserID,
		&post.Title,
		&post.Content,
		pq.Array(&post.Tags),
		&post.Version,
		&post.CratedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, utils.ErrorNotFound
		default:
			return nil, err
		}
	}

	return &post, nil
}
