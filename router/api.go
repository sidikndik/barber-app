package router

import (
	"barber-app/infra"
	"barber-app/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetupReouter(ctx infra.Context) {
	router := gin.Default()

	v1 := router.Group("/v1")
	v1.Use(middleware.Logger())

	// Authentication routes
	auth := v1.Group("/auth")
	{
		auth.POST("/login", ctx.Handler.CustomerHandler.Login)
	}

	// Customer routes
	customer := v1.Group("/customer")
	{
		customer.Use(middleware.Authentication())
		customer.GET("/", ctx.Handler.CustomerHandler.GetAll)
		customer.POST("/", ctx.Handler.CustomerHandler.Create)
	}

	// Product routes
	product := v1.Group("/product")
	{
		product.Use(middleware.Authentication())
		product.GET("/", ctx.Handler.ProductHadler.GetAll)
		product.POST("/", ctx.Handler.ProductHadler.Create)
	}

	fmt.Println("server start on port ", ctx.Config.Port)
	router.Run(":" + ctx.Config.Port)
}
