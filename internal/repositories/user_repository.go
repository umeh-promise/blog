package repositories

import (
	"context"
	"database/sql"

	"github.com/umeh-promise/blog/internal/interfaces"
	"github.com/umeh-promise/blog/internal/models"
	"github.com/umeh-promise/blog/internal/utils"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) interfaces.Users {
	return &UserRepo{db: db}
}

func (repo *UserRepo) Create(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users
			(email, first_name, last_name, username, password, profile_image, role_id) 
		VALUES 
			($1, $2, $3, $4, $5, $6, (SELECT level FROM roles WHERE name=$7)) 
		RETURNING id, created_at, updated_at
	`

	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	role := user.Role
	if role == "" {
		role = "user"
	}

	err := repo.db.QueryRowContext(ctx, query,
		user.Email,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Password.Hash,
		user.ProfileImage,
		role,
	).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return utils.ErrorDuplicateEmail
		case err.Error() == `pq: duplicate key value violates unique constraint "users_username_key"`:
			return utils.ErrorDuplicateUsername
		default:
			return err
		}
	}

	return nil
}

func (repo *UserRepo) GetByID(ctx context.Context, id int64) (*models.User, error) {
	var user models.User

	query := `
		SELECT users.id, email, first_name, last_name, username, password, profile_image, version, created_at, updated_at, roles.name FROM users
		JOIN roles ON (users.role_id = roles.id)
		WHERE users.id=$1
	`

	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	err := repo.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Password.Hash,
		&user.ProfileImage,
		&user.Version,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Role,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, utils.ErrorNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (repo *UserRepo) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	query := `
		SELECT id, email, first_name, last_name, username, password, profile_image, version, created_at, updated_at FROM users
		WHERE email=$1
	`

	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	err := repo.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Password.Hash,
		&user.ProfileImage,
		&user.Version,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, utils.ErrorNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (repo *UserRepo) Update(ctx context.Context, user *models.User) error {
	query := `
		UPDATE users 
		SET first_name = $1, last_name = $2, username = $3, profile_image = $4, version = version + 1
		WHERE id = $5 AND version = $6 
		RETURNING version
	`

	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	err := repo.db.QueryRowContext(ctx, query, user.FirstName, user.LastName, user.Username, user.ProfileImage, user.ID, user.Version).Scan(&user.Version)

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

func (repo *UserRepo) Delete(ctx context.Context, id int64) error {
	query := `
		DELETE FROM users 
		WHERE id = $1	
	`
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	rows, err := repo.db.ExecContext(ctx, query, id)
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
