package controllers

import (
	"bufio"
	"context"
	"fmt"
	"miniproject-nehemia/services"
	"os"
	"strconv"
	"strings"
)

type AdminController struct {
	AdminService *services.AdminService
}

func NewAdminController(service *services.AdminService) *AdminController {
	return &AdminController{AdminService: service}
}

func (ac *AdminController) ShowUsers() {
	users, err := ac.AdminService.GetAllUsers(context.Background())
	if err != nil {
		fmt.Println("Gagal mengambil data user:", err)
		return
	}

	fmt.Println("\n=== DAFTAR USER ===")
	for _, u := range users {
		fmt.Printf("ID: %d | Email: %s | RoleID: %d\n", u.ID, u.Email, u.RoleID)
	}
}

func (ac *AdminController) SetUserRole() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan User ID: ")
	idInput, _ := reader.ReadString('\n')
	userID, _ := strconv.Atoi(strings.TrimSpace(idInput))

	fmt.Print("Masukkan Role Baru (admin/user): ")
	role, _ := reader.ReadString('\n')
	role = strings.TrimSpace(role)

	err := ac.AdminService.UpdateRole(context.Background(), userID, role)
	if err != nil {
		fmt.Println("Gagal mengubah role:", err)
		return
	}

	fmt.Println("Role user berhasil diubah!")
}

func (ac *AdminController) DeleteUser() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan User ID: ")
	idInput, _ := reader.ReadString('\n')
	userID, _ := strconv.Atoi(strings.TrimSpace(idInput))

	err := ac.AdminService.DeleteUser(context.Background(), userID)
	if err != nil {
		fmt.Println("Gagal menghapus user:", err)
		return
	}

	fmt.Println("User berhasil dihapus.")
}
