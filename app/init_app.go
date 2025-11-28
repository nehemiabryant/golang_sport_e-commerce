package app

import (
	"miniproject-nehemia/controllers"
	"miniproject-nehemia/repositories"
	"miniproject-nehemia/services"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Application struct {
	AuthController     *controllers.AuthController
	CategoryController *controllers.CategoryController
	CartController     *controllers.CartController
	PaymentController  *controllers.PaymentController
	AdminController    *controllers.AdminController
	ProductController  *controllers.ProductController
}

func InitApp(db *pgxpool.Pool) *Application {
	// ============ REPOSITORIES ============
	userRepo := &repositories.UserRepository{}
	categoryRepo := repositories.NewCategoryRepository()
	cartRepo := repositories.NewCartRepository()
	productRepo := repositories.NewProductRepository(db)
	paymentRepo := repositories.NewPaymentRepository(db)
	adminRepo := repositories.NewAdminRepository()

	// ============ SERVICES ============
	authService := services.NewAuthService(userRepo)
	categoryService := services.NewCategoryService(categoryRepo)
	cartService := services.NewCartService(cartRepo, productRepo)
	productService := services.NewProductService(productRepo)
	paymentService := services.NewPaymentService(paymentRepo, cartRepo, productRepo)
	adminService := services.NewAdminService(adminRepo)

	// ============ CONTROLLERS ============
	authController := controllers.NewAuthController(authService)
	categoryController := controllers.NewCategoryController(categoryService)
	cartController := controllers.NewCartController(cartService)
	productController := controllers.NewProductController(productService, categoryService)
	paymentController := controllers.NewPaymentController(paymentService)
	adminController := controllers.NewAdminController(adminService)

	// ============ RETURN APP ============
	return &Application{
		AuthController:     authController,
		CategoryController: categoryController,
		CartController:     cartController,
		PaymentController:  paymentController,
		AdminController:    adminController,
		ProductController:  productController,
	}
}
