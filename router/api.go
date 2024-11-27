package router

import (
	"barber-app/infra"
	"barber-app/middleware"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetupReouter(ctx infra.Context) {
	router := gin.New()

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
		product.GET("/", ctx.Handler.ProductHandler.GetAll)
		product.POST("/", ctx.Handler.ProductHandler.Create)
	}

	// Initialize HTTP server
	srv := &http.Server{
		Addr:    ":" + ctx.Config.Port,
		Handler: router,
	}

	// Start the server in a goroutine
	go func() {
		ctx.Log.Info("Starting server on port " + ctx.Config.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			ctx.Log.Fatal("ListenAndServe error: %v", zap.Error(err))
		}
	}()

	// Channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	<-quit
	ctx.Log.Info("Shutdown Server ...")

	// Create context with timeout for graceful shutdown
	ctxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Gracefully shut down the server
	if err := srv.Shutdown(ctxt); err != nil {
		ctx.Log.Fatal("Server Shutdown Error", zap.Error(err))
	}

	ctx.Log.Info("Server exiting")
}
