package repositories

import (
	"context"
	"miniproject-nehemia/config"
	"miniproject-nehemia/models"
)

type AdminRepository struct{}

func NewAdminRepository() *AdminRepository {
	return &AdminRepository{}
}

// AMBIL SEMUA USER
func (ar *AdminRepository) GetAllUsers(ctx context.Context) ([]models.User, error) {
	rows, err := config.DB.Query(ctx, `
		SELECT id, email, password, role_id, created_at
		FROM users
		ORDER BY id ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.ID, &u.Email, &u.Password, &u.RoleID, &u.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

// UPDATE ROLE USER
func (ar *AdminRepository) UpdateUserRole(ctx context.Context, userID int, roleID int) error {
	_, err := config.DB.Exec(ctx,
		`UPDATE users SET role_id = $1 WHERE id = $2`,
		roleID, userID,
	)
	return err
}

// HAPUS USER
func (ar *AdminRepository) DeleteUser(ctx context.Context, userID int) error {
	_, err := config.DB.Exec(ctx,
		`DELETE FROM users WHERE id = $1`,
		userID,
	)
	return err
}
