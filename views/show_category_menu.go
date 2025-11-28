package views

import (
	"bufio"
	"fmt"
	"miniproject-nehemia/controllers"
	"os"
	"strconv"
	"strings"
)

func ShowCategoryMenu(
	categoryController *controllers.CategoryController,
	cartController *controllers.CartController,
	paymentController *controllers.PaymentController,
	userID int,
) {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== KATEGORI OLAHRAGA ===")
		categories, err := categoryController.GetAllCategories()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		for i, c := range categories {
			fmt.Printf("%d. %s\n", i+1, c.Name)
		}

		fmt.Println("0. Logout")
		fmt.Println("99. Lihat Keranjang")
		fmt.Print("Pilih kategori: ")

		input, _ := reader.ReadString('\n')
		choice, _ := strconv.Atoi(strings.TrimSpace(input))

		if choice == 0 {
			fmt.Println("Logout...")
			return
		}

		if choice == 99 {
			ShowCartMenu(cartController, paymentController, func(userID int) {
				ShowCategoryMenu(categoryController, cartController, paymentController, userID)
			}, userID)
			continue
		}

		if choice < 1 || choice > len(categories) {
			fmt.Println("Pilihan tidak valid")
			continue
		}

		selected := categories[choice-1]
		ShowProductMenu(categoryController, cartController, userID, selected.ID)
	}
}
