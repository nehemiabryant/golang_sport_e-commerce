package controllers

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"miniproject-nehemia/services"
)

type ProductController struct {
	productService  *services.ProductService
	categoryService *services.CategoryService
}

func NewProductController(ps *services.ProductService, cs *services.CategoryService) *ProductController {
	return &ProductController{
		productService:  ps,
		categoryService: cs,
	}
}

func readInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// ADD NEW PRODUCT - ADMIN ONLY
func (c *ProductController) AddProductCLI() {
	ctx := context.Background()

	id := readInput("ID Produk: ")
	name := readInput("Nama produk: ")
	priceStr := readInput("Harga produk: ")
	stockStr := readInput("Stok awal: ")

	price, _ := strconv.Atoi(priceStr)
	stock, _ := strconv.Atoi(stockStr)

	// Ambil kategori
	ids, names, err := c.categoryService.GetCategories(ctx)
	if err != nil {
		fmt.Println("Gagal mengambil kategori:", err)
		return
	}

	// Tampilkan kategori
	fmt.Println("\n=== PILIH KATEGORI PRODUK ===")
	for i := range ids {
		fmt.Printf("%d. %s (ID: %s)\n", i+1, names[i], ids[i])
	}

	input := readInput("Pilih kategori (nomor): ")
	idx, _ := strconv.Atoi(input)
	idx--

	if idx < 0 || idx >= len(ids) {
		fmt.Println("Kategori tidak valid!")
		return
	}

	selectedCategory := ids[idx]

	// Tambah produk
	err = c.productService.AddProduct(ctx, id, name, float64(price), stock)
	if err != nil {
		fmt.Println("Gagal menambahkan produk:", err)
		return
	}

	// Assign to category
	err = c.productService.AssignProductToCategory(ctx, id, selectedCategory)
	if err != nil {
		fmt.Println("Gagal memasukkan produk ke kategori:", err)
		return
	}

	fmt.Println("\nProduk berhasil ditambahkan dan masuk ke kategori!")
}

// RESTOCK PRODUCT
func (c *ProductController) RestockProductCLI() {
	id := readInput("ID Produk: ")
	qtyStr := readInput("Jumlah restock: ")

	qty, _ := strconv.Atoi(qtyStr)

	err := c.productService.RestockProduct(context.Background(), id, qty)
	if err != nil {
		fmt.Println("Gagal restock:", err)
		return
	}

	fmt.Println("Restock berhasil!")
}

// DECREASE STOCK
func (c *ProductController) DecreaseStockCLI() {
	id := readInput("ID Produk: ")
	qtyStr := readInput("Jumlah pengurangan stok: ")

	qty, _ := strconv.Atoi(qtyStr)

	err := c.productService.DecreaseStock(context.Background(), id, qty)
	if err != nil {
		fmt.Println("Gagal mengurangi stok:", err)
		return
	}

	fmt.Println("Stok berhasil dikurangi!")
}

// LIST ALL PRODUCTS
func (c *ProductController) ListAllProductsCLI() {
	ctx := context.Background()

	grouped, err := c.productService.GetProductsGroupedByCategory(ctx)
	if err != nil {
		fmt.Println("Gagal mengambil daftar produk:", err)
		return
	}

	fmt.Println("\n=========== DAFTAR PRODUK PER KATEGORI ===========")

	for category, products := range grouped {
		fmt.Printf("\n--- %s ---\n", strings.ToUpper(category))

		if len(products) == 0 {
			fmt.Println("(Tidak ada produk)")
			continue
		}

		for _, p := range products {
			fmt.Printf("ID: %s | Nama: %s | Harga: %.2f | Stok: %d\n",
				p.ID, p.Name, p.Price, p.Stock)
		}
	}

	fmt.Println("===================================================\n")
}
