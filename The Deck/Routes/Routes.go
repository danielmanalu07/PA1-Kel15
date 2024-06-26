package routes

import (
	controllers "api/the_deck/Controllers"
	middleware "api/the_deck/Middleware"
	service "api/the_deck/Service"

	"github.com/gofiber/fiber/v2"
)

func RouteAdmin(App *fiber.App, adminController *controllers.AdminController) {
	admin := App.Group("/admin")
	admin.Post("/login", adminController.AdminLogin)
	admin.Use(middleware.CheckLogin())
	admin.Get("/profile", adminController.GetProfile)
	admin.Post("/logout", adminController.LogoutAdmin)
	admin.Put("/order/:id", adminController.UpdateStatusOrder)
	admin.Put("/table/:id", adminController.ApproveReqTable)
	
}

func RouteCategory(App *fiber.App, categoryController *controllers.CategoryController) {
	category := App.Group("/category")
	category.Get("/", categoryController.CategoryGetAll)
	category.Use(middleware.CheckLogin())
	category.Post("/create", categoryController.CategoryCreate)
	category.Get("/:id", categoryController.CategoryGetById)
	category.Put("/edit/:id", categoryController.CategoryUpdate)
	category.Delete("/delete/:id", categoryController.CategoryDelete)
}

func RouteProduct(App *fiber.App, productController *controllers.ProductController) {
	product := App.Group("/product")
	product.Static("/image", service.PathImageProduct)
	product.Get("/", productController.ProductGetAll)
	product.Get("/category/:cat", productController.ProductGetByCategory)
	product.Get("/:id", productController.ProductGetById)
	product.Use(middleware.CheckLogin())
	product.Post("/create", productController.ProductCreate)
	product.Put("/edit/:id", productController.ProductUpdate)
	product.Delete("/delete/:id", productController.ProductDelete)
}

func RouteTable(App *fiber.App, tableController *controllers.TableController) {
	table := App.Group("/table")
	table.Get("/", tableController.TableGetAll)
	table.Use(middleware.CheckLogin())
	table.Post("/create", tableController.TableCreate)
	table.Get("/:id", tableController.TableGetById)
	table.Put("/edit/:id", tableController.TableUpdate)
	table.Delete("/delete/:id", tableController.TableDelete)
}

func RouteCustomer(App *fiber.App, customerController *controllers.CustomerController) {
	customer := App.Group("/customer")
	customer.Post("/register", customerController.CustomerRegister)
	customer.Post("/login", customerController.CustomerLogin)
	customer.Static("/image", service.PathImageCustomer)
	customer.Put("/forgot-password", customerController.CustomerForgotPassword)
	customer.Use(middleware.CheckCustomer())
	customer.Get("/profile", customerController.GetProfile)
	customer.Post("/logout", customerController.CustomerLogout)
	customer.Put("/update-profile", customerController.CustomerUpdateProfile)
	customer.Put("/edit-password", customerController.CustomerEditPassword)
}

func RouteCart(App *fiber.App, cartController *controllers.CartController) {
	cart := App.Group("/cart")
	cart.Use(middleware.CheckCustomer())
	cart.Post("/add", cartController.AddItemCart)
	cart.Get("/myCart", cartController.GetItemMyCart)
	cart.Delete("/delete/:id", cartController.DeleteMyCart)
	cart.Put("/edit/:id", cartController.UpdateQuantity)
}

func RouteOrder(App *fiber.App, orderController *controllers.OrderController) {
	order := App.Group("/order")
	order.Get("/", orderController.GetAllOrder)
	order.Static("/image", service.PathImageOrder)
	order.Use(middleware.CheckCustomer())
	order.Put("/status/:id", orderController.UpdateStatus)
	order.Post("/create", orderController.CustomerCreateOrder)
	order.Get("/myorder", orderController.GetMyOrder)
	order.Put("/payment/:id", orderController.ProofOfPayment)
}

func RouteRequestTable(App *fiber.App, requestController *controllers.RequestTableController) {
	rt := App.Group("/requestTable")
	rt.Get("/", requestController.GetAllRequest)
	rt.Use(middleware.CheckCustomer())
	rt.Post("/create", requestController.CreateRequestTable)
	rt.Get("/myRequest", requestController.GetMyReqTable)
	rt.Put("/status/:id", requestController.UpdateStatus)
}
