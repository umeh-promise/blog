package repositories

import (
	"context"
	"database/sql"

	"github.com/umeh-promise/blog/internal/interfaces"
	"github.com/umeh-promise/blog/internal/models"
	"github.com/umeh-promise/blog/internal/utils"
)

type RoleRepository struct {
	db *sql.DB
}

func NewRoleRepository(db *sql.DB) interfaces.Role {
	return &RoleRepository{db: db}
}

func (repo *RoleRepository) GetByName(ctx context.Context, name string) (*models.Role, error) {
	var role models.Role
	query := `
		SELECT id, name, level, description FROM roles
		WHERE name = $1
	`

	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	err := repo.db.QueryRowContext(ctx, query, name).Scan(
		&role.ID,
		&role.Name,
		&role.Level,
		&role.Description,
	)
	if err != nil {
		return nil, err
	}

	return &role, nil
}
