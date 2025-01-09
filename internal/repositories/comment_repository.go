package repositories

import (
	"context"
	"database/sql"

	"github.com/umeh-promise/blog/internal/interfaces"
	"github.com/umeh-promise/blog/internal/models"
	"github.com/umeh-promise/blog/internal/utils"
)

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) interfaces.Comment {
	return &CommentRepository{db: db}
}

func (repo *CommentRepository) Create(ctx context.Context, comment *models.Comment) error {
	query := `
		INSERT INTO comments (user_id, post_id, content)
		VALUES ($1, $2, $3) RETURNING id, created_at, updated_at
	`

	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	err := repo.db.QueryRowContext(ctx, query,
		comment.UserID,
		comment.PostID,
		comment.Content,
	).Scan(
		&comment.ID,
		&comment.CreatedAt,
		&comment.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (repo *CommentRepository) GetByPostID(ctx context.Context, id int64) ([]models.Comment, error) {
	comments := []models.Comment{}

	query := `
		SELECT c.id, c.user_id, c.post_id, c.content, c.created_at, c.updated_at, users.id, users.username, users.first_name, users.last_name FROM comments c
		JOIN users ON users.id = c.user_id
		WHERE c.post_id = $1
		ORDER BY c.created_at
		`

	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	rows, err := repo.db.QueryContext(ctx, query, id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, utils.ErrorNotFound
		default:
			return nil, err
		}
	}
	defer rows.Close()

	for rows.Next() {
		var comment models.Comment
		comment.User = models.User{}
		err := rows.Scan(
			&comment.ID,
			&comment.UserID,
			&comment.PostID,
			&comment.Content,
			&comment.CreatedAt,
			&comment.UpdatedAt,
			&comment.User.ID,
			&comment.User.Username,
			&comment.User.FirstName,
			&comment.User.LastName,
		)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)

	}

	return comments, nil
}
