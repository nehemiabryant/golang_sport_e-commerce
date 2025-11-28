package controllers

import (
	"context"
	"fmt"
	"miniproject-nehemia/helper"
	"miniproject-nehemia/models"
	"miniproject-nehemia/services"
)

type CartController struct {
	service *services.CartService
}

func NewCartController(service *services.CartService) *CartController {
	return &CartController{service: service}
}

func (cc *CartController) AddToCartDirect(userID int, productID string, qty int) error {
	return cc.service.AddToCart(context.Background(), userID, productID, qty)
}

func (cc *CartController) AddToCartCLI(userID int) {
	var productID string
	var qty int

	fmt.Print("Masukkan Product ID (CHAR4): ")
	fmt.Scanln(&productID)

	fmt.Print("Jumlah: ")
	fmt.Scanln(&qty)

	err := cc.service.AddToCart(context.Background(), userID, productID, qty)
	if err != nil {
		fmt.Println("Gagal menambah ke keranjang:", err)
		return
	}

	fmt.Println("Produk berhasil ditambahkan ke keranjang!")
}

func (cc *CartController) ShowCartTable(userID int) ([]models.CartItemDetail, float64, error) {
	items, err := cc.service.GetCartItems(context.Background(), userID)
	if err != nil {
		return nil, 0, err
	}

	var total float64
	for _, item := range items {
		total += item.Subtotal
	}

	return items, total, nil
}

func (cc *CartController) ShowCartCLI(userID int) {
	items, err := cc.service.GetCartItems(context.Background(), userID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\n================= KERANJANG BELANJA =================")
	if len(items) == 0 {
		fmt.Println("(Kosong)")
		return
	}

	fmt.Println("--------------------------------------------------------------------------")
	fmt.Printf("| %-6s | %-25s | %-5s | %-10s | %-10s |\n",
		"ID", "Nama Produk", "Qty", "Harga", "Subtotal")
	fmt.Println("--------------------------------------------------------------------------")

	var total float64
	for _, item := range items {
		fmt.Printf(
			"| %-6s | %-25s | %-6d | %-15s | %-15s |\n",
			item.ProductID,
			item.Name,
			item.Quantity,
			helper.FormatRupiah(item.Price),
			helper.FormatRupiah(item.Subtotal),
		)
		total += item.Subtotal
	}

	fmt.Println("--------------------------------------------------------------------------")
	fmt.Printf("Total: %s\n", helper.FormatRupiah(total))
}

// =================== TAMBAH QTY ===================
func (cc *CartController) AddQty(userID int, productID string) error {
	return cc.service.AddQty(context.Background(), userID, productID)
}

// =================== KURANGI QTY ===================
func (cc *CartController) ReduceQty(userID int, productID string) error {
	return cc.service.ReduceQty(context.Background(), userID, productID)
}

// =================== HAPUS ITEM ===================
func (cc *CartController) RemoveItemCLI(userID int, productID string) error {
	return cc.service.RemoveItem(context.Background(), userID, productID)
}
