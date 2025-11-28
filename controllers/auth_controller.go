package controllers

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"syscall"

	"miniproject-nehemia/helper"
	"miniproject-nehemia/services"

	"golang.org/x/term"
)

type AuthController struct {
	AuthService *services.AuthService
}

func NewAuthController(as *services.AuthService) *AuthController {
	return &AuthController{
		AuthService: as,
	}
}

// Helper untuk membaca input string biasa (misalnya email)
func readStringInput(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Helper untuk membaca password tersembunyi
func readHiddenPassword(prompt string) (string, error) {
	fmt.Print(prompt)
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()

	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(bytePassword)), nil
}

func (c *AuthController) SignupCLI() {
	reader := bufio.NewReader(os.Stdin)

	for {
		email := readStringInput(reader, "Email: ")

		if !helper.IsValidEmail(email) {
			fmt.Println("Format email tidak valid.")
			continue
		}

		password, err := readHiddenPassword("Password: ")
		if err != nil {
			fmt.Println("Gagal membaca password:", err)
			continue
		}

		if !helper.IsValidPassword(password) {
			fmt.Println("Password minimal 8 karakter dan harus mengandung huruf & angka.")
			continue
		}

		err = c.AuthService.Register(context.Background(), email, password)
		if err != nil {
			fmt.Println("Signup gagal:", err)
			continue
		}

		fmt.Println("Signup berhasil!")
		break
	}
}

func (c *AuthController) LoginCLI() (int, string) {
	reader := bufio.NewReader(os.Stdin)

	email := readStringInput(reader, "Email: ")
	password, _ := readHiddenPassword("Password: ")

	id, RoleID, err := c.AuthService.Login(context.Background(), email, password)
	if err != nil {
		fmt.Println("Login gagal:", err)
		return 0, ""
	}

	fmt.Println("Login berhasil!", RoleID)
	return id, RoleID
}
