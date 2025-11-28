package services

import (
	"context"
	"errors"
	"miniproject-nehemia/models"
	"miniproject-nehemia/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

func NewAuthService(repo *repositories.UserRepository) *AuthService {
	return &AuthService{
		userRepo: repo,
	}
}

func (as *AuthService) Register(ctx context.Context, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Jika email admin, set role_id = 2
	roleID := 1
	if email == "admin@gmail.com" {
		roleID = 2
	}

	return as.userRepo.CreateUser(ctx, email, string(hashedPassword), roleID)
}

// Login sekarang mengembalikan userID
func (as *AuthService) Login(ctx context.Context, email, password string) (int, string, error) {
	user, err := as.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return 0, "", errors.New("email atau password salah")
	}

	// cek password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return 0, "", errors.New("email atau password salah")
	}

	// convert roleID → roleName
	roleName := models.GetRoleName(user.RoleID)

	return user.ID, roleName, nil
}

func (s *AuthService) CreateAdmin(ctx context.Context, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.userRepo.CreateAdminUser(ctx, email, string(hashedPassword))
}
