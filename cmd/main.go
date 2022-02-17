package main

import (
	"context"

	"github.com/ATM/pkg/handlers"
	"github.com/ATM/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	ctx := context.Background()
	atmHandler := handlers.NewATMHandler(ctx)

	authorized := r.Group("/")
	authorized.Use(middleware.APIKeyRequired())
	{
		authorized.POST("/login", atmHandler.LoginWithPin)
	}

	transaction := r.Group("/")
	transaction.Use(middleware.TokenRequired())
	{
		transaction.GET("/balance", atmHandler.GetBalance)
	}

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
