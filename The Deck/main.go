package main

import (
	database "api/the_deck/Database"
	migrations "api/the_deck/Database/Migrations"
	routes "api/the_deck/Routes"
	settings "api/the_deck/Settings"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connection()
	migrations.Migration()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "https://gofiber.io",
	}))

	adminController := settings.SetUpServiceAdmin()
	categoryController := settings.SetUpServiceCategory()
	productController := settings.SetUpServiceProduct()
	tableController := settings.SetUpServiceTable()
	customerController := settings.SetUpServiceCustomer()
	cartController := settings.SetUpServiceCart()
	orderController := settings.SetUpServiceOrder()
	requestTableController := settings.SetUpServiceRequestTable()

	routes.RouteAdmin(app, adminController)
	routes.RouteCategory(app, categoryController)
	routes.RouteProduct(app, productController)
	routes.RouteTable(app, tableController)
	routes.RouteCustomer(app, customerController)
	routes.RouteCart(app, cartController)
	routes.RouteOrder(app, orderController)
	routes.RouteRequestTable(app, requestTableController)

<<<<<<< HEAD
	err := app.Listen("192.168.188.215:8080")
=======
	err := app.Listen("172.27.1.162:8080")
>>>>>>> c5757912099616d39fa31b5716dec2f30ff465cd

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
}
