package services

import (
	"context"
	"errors"
	"miniproject-nehemia/models"
	"miniproject-nehemia/repositories"
)

type AdminService struct {
	Repo *repositories.AdminRepository
}

func NewAdminService(repo *repositories.AdminRepository) *AdminService {
	return &AdminService{Repo: repo}
}

// Show all users
func (s *AdminService) GetAllUsers(ctx context.Context) ([]models.User, error) {
	return s.Repo.GetAllUsers(ctx)
}

// Set role
func (s *AdminService) UpdateRole(ctx context.Context, userID int, newRole string) error {
	var roleID int

	if newRole == "admin" {
		roleID = 2
	} else if newRole == "user" {
		roleID = 1
	} else {
		return errors.New("role harus admin atau user")
	}

	return s.Repo.UpdateUserRole(ctx, userID, roleID)
}

// Delete user
func (s *AdminService) DeleteUser(ctx context.Context, userID int) error {
	return s.Repo.DeleteUser(ctx, userID)
}
