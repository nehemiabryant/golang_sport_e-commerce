package views

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"miniproject-nehemia/controllers"
	"miniproject-nehemia/helper"
)

func ShowProductMenu(categoryController *controllers.CategoryController, cartController *controllers.CartController, userID int, categoryID string) {
	reader := bufio.NewReader(os.Stdin)

	for {
		products, err := categoryController.GetProductsByCategory(categoryID)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("\n=== PRODUK DALAM KATEGORI %s ===\n", categoryID)
		fmt.Println("--------------------------------------------------------------")
		fmt.Printf("| %-3s | %-25s | %-15s | %-5s |\n", "No", "Nama Produk", "Harga", "Stok")
		fmt.Println("--------------------------------------------------------------")

		for i, p := range products {
			fmt.Printf(
				"| %-3d | %-25s | %-15s | %-5d |\n",
				i+1, p.Name, helper.FormatRupiah(p.Price), p.Stock,
			)
		}

		fmt.Println("--------------------------------------------------------------")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih produk untuk tambah ke keranjang: ")

		input, _ := reader.ReadString('\n')
		choice, _ := strconv.Atoi(strings.TrimSpace(input))

		if choice == 0 {
			return
		}

		if choice < 1 || choice > len(products) {
			fmt.Println("Pilihan tidak valid.")
			continue
		}

		selected := products[choice-1]

		fmt.Printf("Masukkan jumlah untuk %s: ", selected.Name)
		qtyInput, _ := reader.ReadString('\n')
		qty, _ := strconv.Atoi(strings.TrimSpace(qtyInput))

		if qty <= 0 {
			fmt.Println("Jumlah tidak boleh kurang dari 1!")
			continue
		}

		err = cartController.AddToCartDirect(userID, selected.ID, qty)
		if err != nil {
			fmt.Println("Gagal menambah ke keranjang:", err)
			continue
		}

		fmt.Println("Produk berhasil ditambahkan ke keranjang!")
	}
}
