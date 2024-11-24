package api

import (
	"backend/internal/config"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"

	billApi "backend/internal/api/bill"
	transactionApi "backend/internal/api/transaction"
	userApi "backend/internal/api/user"

	middleware "backend/internal/api/middleware"

	_ "backend/docs"
)

// @title NashCash API
// @version 1.0
// @description Бакенд всего этого чуда. In Process
// @BasePath /
func Serve(conf config.ApiConfig) {
	portString := strconv.Itoa(conf.Port)
	app := fiber.New()

	// Swagger setup
	app.Get("/docs/*", swagger.HandlerDefault) // default
	app.Get("/docs/*", swagger.New(swagger.Config{
		URL: "http://localhost" + portString + "/doc.json",
	}))

	api := app.Group("/api", logger.New())

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", userApi.Login)

	// User
	user := api.Group("/user")
	user.Post("/", userApi.CreateUser)
	user.Get("/:id", middleware.Protected(), userApi.GetUser)
	user.Patch("/:id", middleware.Protected(), userApi.UpdateUser)
	user.Delete("/:id", middleware.Protected(), userApi.DeleteUser)

	// Bill
	bill := api.Group("/bill")
	bill.Post("/", middleware.Protected(), billApi.CreateBill)
	bill.Get("/all", middleware.Protected(), billApi.GetAllUserBills)
	bill.Get("/:id", middleware.Protected(), billApi.GetBill)
	bill.Delete("/:id", middleware.Protected(), billApi.DeleteBill)
	bill.Put("/:id", middleware.Protected(), billApi.UpdateBill)

	// transaction
	transactions := api.Group("/transaction")
	transactions.Post("/", middleware.Protected(), transactionApi.CreateTransaction)
	transactions.Patch("/accept/:id", middleware.Protected(), transactionApi.AcceptTransaction)
	transactions.Patch("/decline/:id", middleware.Protected(), transactionApi.DeclineTransaction)
	transactions.Patch("/resolve/:id", middleware.Protected(), transactionApi.ResolveTransaction)
	transactions.Put("/:id", middleware.Protected(), transactionApi.UpdateTransaction)
	transactions.Delete("/:id", middleware.Protected(), transactionApi.DeleteTransaction)

	app.Listen(":" + portString)
}
