package repositories

import (
	"context"
	"miniproject-nehemia/config"
	"miniproject-nehemia/models"
)

type UserRepository struct{}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT id, email, password, role_id, created_at FROM users WHERE email = $1`

	row := config.DB.QueryRow(ctx, query, email)

	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.RoleID, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, email, password string, roleID int) error {
	_, err := config.DB.Exec(ctx,
		`INSERT INTO users (email, password, role_id) VALUES ($1, $2, $3)`,
		email, password, roleID,
	)
	return err
}

func (ur *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	row := config.DB.QueryRow(ctx,
		`SELECT id, email, password, role_id, created_at FROM users WHERE email = $1`,
		email,
	)

	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.RoleID, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) CreateAdminUser(ctx context.Context, email, password string) error {
	_, err := config.DB.Exec(ctx,
		`INSERT INTO users (email, password, role_id) VALUES ($1, $2, 2)`, // 2 = admin
		email, password,
	)
	return err
}
