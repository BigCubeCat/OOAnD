package api

import (
	"backend/internal/config"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"

	billApi "backend/internal/api/bill"
	handler "backend/internal/api/handlers"
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
	auth.Post("/login", handler.Login)

	// User
	user := api.Group("/user")
	user.Get("/:id", handler.GetUser)
	user.Post("/", handler.CreateUser)
	user.Patch("/:id", middleware.Protected(), handler.UpdateUser)
	user.Delete("/:id", middleware.Protected(), handler.DeleteUser)

	// Bill
	bill := api.Group("/bill")
	bill.Post("/", billApi.CreateBill)
	bill.Get("/all", billApi.GetAllUserBills)
	bill.Get("/:id", billApi.GetBill)

	app.Listen(":" + portString)
}
