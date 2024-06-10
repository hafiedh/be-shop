package app

import (
	"be-shop/internal/app/controller"
	"be-shop/pkg/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func setRoute(
	e *echo.Echo,

	authCtrl controller.AuthCtrl,
	productCtrl controller.ProductCtrl,
	middleware middleware.MiddleWare,
) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	base := e.Group("/v1")

	users := base.Group("/users")
	{
		users.POST("/register", authCtrl.UserRegistration)
		users.POST("/login", authCtrl.UserLogin)
	}

	products := base.Group("/products")
	{
		products.GET("", productCtrl.GetAllProduct)
		products.POST("", productCtrl.CreateProduct)
		products.GET("/:id", productCtrl.GetProductByID)
		products.PATCH("/:id", productCtrl.UpdateProductPrice)
		products.DELETE("/:id", productCtrl.DeleteProduct)
		products.GET("/category/:id", productCtrl.GetProductsByCategoryID)
	}

	base.Use(middleware.AuthUser)
}
