package repositories

import (
	"context"
	"database/sql"
	"errors"

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
	query := `
		INSERT INTO posts (user_id, title, content, tags) 
		VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`

	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	err := postRepo.db.QueryRowContext(ctx, query, post.UserID, post.Title, post.Content, pq.Array(post.Tags)).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (postRepo *PostRepository) GetByID(ctx context.Context, id int64) (*models.Post, error) {

	var post models.Post

	query := `SELECT id, user_id, title, content, tags, version, created_at, updated_at FROM posts WHERE id=$1`

	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	err := postRepo.db.QueryRowContext(ctx, query, id).Scan(
		&post.ID,
		&post.UserID,
		&post.Title,
		&post.Content,
		pq.Array(&post.Tags),
		&post.Version,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, utils.ErrorNotFound
		default:
			return nil, err
		}
	}

	return &post, nil
}

func (repo *PostRepository) GetAll(ctx context.Context) ([]models.Post, error) {
	query := `
		SELECT id, user_id, title, content, tags, version, created_at, updated_at FROM posts
	`

	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		post := models.Post{}
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, pq.Array(&post.Tags), &post.Version, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (postRepo *PostRepository) Update(ctx context.Context, post *models.Post) error {
	query := `
		UPDATE posts 
		SET title = $1, content = $2, tags = $3, version = version + 1
		WHERE id = $4 AND version = $5
		RETURNING version
	`
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	err := postRepo.db.QueryRowContext(ctx, query, post.Title, post.Content, pq.Array(post.Tags), post.ID, post.Version).Scan(&post.Version)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return utils.ErrorNotFound
		default:
			return err
		}
	}
	return nil
}

func (postRepo *PostRepository) Delete(ctx context.Context, id int64) error {
	query := `
		DELETE FROM posts
		WHERE id = $1	
	`
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	rows, err := postRepo.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	row, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if row <= 0 {
		return utils.ErrorNotFound
	}

	return nil
}
