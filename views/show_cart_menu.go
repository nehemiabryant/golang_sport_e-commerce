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

func ShowCartMenu(
	cartController *controllers.CartController,
	paymentController *controllers.PaymentController,
	categoryMenuFunc func(int),
	userID int,
) {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Ambil cart data tanpa printing
		items, total, err := cartController.ShowCartTable(userID)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("\n================= KERANJANG BELANJA =================")

		if len(items) == 0 {
			fmt.Println("(Kosong)")
		} else {
			fmt.Println("--------------------------------------------------------------------------")
			fmt.Printf("| %-6s | %-25s | %-5s | %-10s | %-10s |\n",
				"ID", "Nama Produk", "Qty", "Harga", "Subtotal")
			fmt.Println("--------------------------------------------------------------------------")

			for _, item := range items {
				fmt.Printf(
					"| %-6s | %-25s | %-6d | %-15s | %-15s |\n",
					item.ProductID,
					item.Name,
					item.Quantity,
					helper.FormatRupiah(item.Price),
					helper.FormatRupiah(item.Subtotal),
				)
			}

			fmt.Println("--------------------------------------------------------------------------")
			fmt.Printf("Total: %s\n", helper.FormatRupiah(total))
		}

		// Menu user
		fmt.Println("\n=== MENU KERANJANG ===")
		fmt.Println("1. Tambah Qty")
		fmt.Println("2. Kurangi Qty")
		fmt.Println("3. Hapus Item")
		fmt.Println("4. Tambah Item Lain (ke menu kategori olahraga)")
		fmt.Println("5. Checkout Pembayaran")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih menu: ")

		input, _ := reader.ReadString('\n')
		choice, _ := strconv.Atoi(strings.TrimSpace(input))

		switch choice {

		case 0:
			return

		case 1, 2, 3:
			fmt.Print("Masukkan Product ID: ")
			productID, _ := reader.ReadString('\n')
			productID = strings.TrimSpace(productID)

			if choice == 1 {
				fmt.Println("Masukan jumlah yang ingin ditambahkan:")
				var qty int
				fmt.Scanln(&qty)
				for i := 0; i < qty; i++ {
					cartController.AddQty(userID, productID)
				}
			} else if choice == 2 {
				fmt.Println("Masukan jumlah yang ingin dikurangi:")
				var qty int
				fmt.Scanln(&qty)
				for i := 0; i < qty; i++ {
					cartController.ReduceQty(userID, productID)
				}
			} else {
				cartController.RemoveItemCLI(userID, productID)
			}

		case 4:
			categoryMenuFunc(userID) // kembali ke kategori

		case 5:
			paymentController.ProcessPaymentCLI(userID)
			return

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
