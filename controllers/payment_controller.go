package controllers

import (
	"bufio"
	"context"
	"fmt"
	"miniproject-nehemia/helper"
	"miniproject-nehemia/services"
	"os"
	"strconv"
	"strings"
)

type PaymentController struct {
	paymentService *services.PaymentService
}

func NewPaymentController(s *services.PaymentService) *PaymentController {
	return &PaymentController{paymentService: s}
}

func (pc *PaymentController) ProcessPaymentCLI(userID int) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n=== PILIH METODE PEMBAYARAN ===")
	fmt.Println("1. Cash")
	fmt.Println("2. Credit Card")
	fmt.Print("Pilihan: ")

	input, _ := reader.ReadString('\n')
	choice, _ := strconv.Atoi(strings.TrimSpace(input))

	switch choice {
	case 1:
		pc.processCashPayment(userID)
	case 2:
		pc.processCreditCardPayment(userID)
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

//	CASH PAYMENT

func (pc *PaymentController) processCashPayment(userID int) {
	reader := bufio.NewReader(os.Stdin)

	total, err := pc.paymentService.GetCartTotal(context.Background(), userID)
	if err != nil {
		fmt.Println("Gagal menghitung total:", err)
		return
	}

	fmt.Println("Total Belanja:", helper.FormatRupiah(total))
	fmt.Print("Masukkan nominal pembayaran: ")

	input, _ := reader.ReadString('\n')
	bayar, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil || bayar <= 0 {
		fmt.Println("Input tidak valid!")
		return
	}

	if bayar < total {
		fmt.Println("Nominal kurang!")
		return
	}

	kembalian := bayar - total

	paymentID, err := pc.paymentService.ProcessPayment(context.Background(), userID)
	if err != nil {
		fmt.Println("Gagal memproses pembayaran:", err)
		return
	}

	fmt.Println("\n=== PEMBAYARAN CASH BERHASIL ===")
	fmt.Println("ID Pembayaran :", paymentID)
	fmt.Println("Total         :", helper.FormatRupiah(total))
	fmt.Println("Dibayar       :", helper.FormatRupiah(bayar))
	fmt.Println("Kembalian     :", helper.FormatRupiah(kembalian))
}

//	Credit Card

func (pc *PaymentController) processCreditCardPayment(userID int) {
	reader := bufio.NewReader(os.Stdin)

	total, err := pc.paymentService.GetCartTotal(context.Background(), userID)
	if err != nil {
		fmt.Println("Gagal menghitung total:", err)
		return
	}
	fmt.Println("Total Belanja:", helper.FormatRupiah(total))
	fmt.Print("Masukkan Tenor (3 bulan, 6 bulan, 12 bulan): ")
	input, _ := reader.ReadString('\n')
	tenor, _ := strconv.Atoi(strings.TrimSpace(input))

	if tenor != 3 && tenor != 6 && tenor != 12 {
		fmt.Println("Tenor tidak valid!")
		return
	}
	monthlyPayment := total / float64(tenor)

	paymentID, err := pc.paymentService.ProcessPayment(context.Background(), userID)
	if err != nil {
		fmt.Println("Gagal memproses pembayaran:", err)
		return
	}
	fmt.Println("\n=== PEMBAYARAN KARTU KREDIT BERHASIL ===")
	fmt.Println("ID Pembayaran     :", paymentID)
	fmt.Println("Total             :", helper.FormatRupiah(total))
	fmt.Printf("Tenor             : %d bulan\n", tenor)
	fmt.Println("Cicilan per bulan : ", helper.FormatRupiah(monthlyPayment))
}

// VIEW ALL PAYMENTS - ADMIN ONLY
func (pc *PaymentController) ShowAllPayments() {
	data, err := pc.paymentService.GetAllPayments(context.Background())
	if err != nil {
		fmt.Println("Gagal mengambil riwayat pembayaran:", err)
		return
	}

	fmt.Println("\n=== RIWAYAT PEMBELIAN SEMUA USER ===")
	for _, d := range data {
		fmt.Printf(
			"PaymentID: %d | User: %s | Produk: %s | Qty: %d | Subtotal: %.2f | Total: %.2f | Tanggal: %s\n",
			d.PaymentID,
			d.UserEmail,
			d.ProductName,
			d.Quantity,
			d.Subtotal,
			d.TotalAmount,
			d.CreatedAt.Format("2006-01-02 15:04"),
		)
	}
	fmt.Println("=====================================")
}
