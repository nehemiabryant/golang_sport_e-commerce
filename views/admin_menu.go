package views

import (
	"bufio"
	"fmt"
	"miniproject-nehemia/app"
	"os"
	"strconv"
	"strings"
)

func ShowAdminMenu(app *app.Application, userID int) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n===== ADMIN MENU =====")
		fmt.Println("1. Lihat Semua User")
		fmt.Println("2. Lihat Semua Produk")
		fmt.Println("3. Lihat Riwayat Pembelian Semua User")
		fmt.Println("0. Logout")
		fmt.Print("Pilih: ")

		input, _ := reader.ReadString('\n')
		choice, _ := strconv.Atoi(strings.TrimSpace(input))

		switch choice {
		case 1:
			app.AdminController.ShowUsers() // Menampilkan semua user
		case 2:
			app.ProductController.ListAllProductsCLI() // Menampilkan semua produk
			showProductSubmenu(app)                    // submenu untuk tambah/restock/kurangi
		case 3:
			app.PaymentController.ShowAllPayments() // Lihat riwayat pembelian semua user
		case 0:
			fmt.Println("Logout...")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func showProductSubmenu(app *app.Application) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== MENU PRODUK ===")
		fmt.Println("1. Tambah Produk")
		fmt.Println("2. Restock Produk")
		fmt.Println("3. Kurangi Stok Produk")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih: ")

		input, _ := reader.ReadString('\n')
		choice, _ := strconv.Atoi(strings.TrimSpace(input))

		switch choice {
		case 1:
			app.ProductController.AddProductCLI()
		case 2:
			app.ProductController.RestockProductCLI()
		case 3:
			app.ProductController.DecreaseStockCLI()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}
