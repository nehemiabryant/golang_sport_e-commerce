package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"miniproject-nehemia/app"
	"miniproject-nehemia/config"
	"miniproject-nehemia/views"
)

func main() {
	if err := config.ConnectDB(); err != nil {
		panic(err)
	}

	// init semua controller lewat 1 file
	db := config.DB
	application := app.InitApp(db)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("=== MENU AUTH ===")
		fmt.Println("1. Signup")
		fmt.Println("2. Login")
		fmt.Println("0. Exit")
		fmt.Print("Masukkan pilihan: ")

		input, _ := reader.ReadString('\n')
		choice, _ := strconv.Atoi(strings.TrimSpace(input))

		switch choice {
		case 0:
			fmt.Println("Keluar...")
			return
		case 1:
			application.AuthController.SignupCLI()
		case 2:
			userID, role := application.AuthController.LoginCLI()
			if userID != 0 {
				if role == "admin" { // admin
					views.ShowAdminMenu(application, userID)
				} else {
					views.ShowCategoryMenu(application.CategoryController, application.CartController, application.PaymentController, userID)
				}
			}
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}
