package cmd

import (
	"authen-system/internal/config"
	"authen-system/internal/controllers"
	"authen-system/internal/database"
	"fmt"
	"github.com/gin-gonic/gin"
)

func registerUserHandler(g gin.IRouter, config config.App) {
	db, err := database.ProvideSQL(config.MySQL)
	if err != nil {
		fmt.Printf("failed to provide SQL err = %v", err)
		return
	}
	userRepo := database.NewUserRepository(db)
	handler := controllers.NewHandler(userRepo, config.Authentication)
	userAPI := g.Group("/user")
	{
		userAPI.POST("/login", handler.Login)
		userAPI.POST("/signup", handler.SignUp)
	}
}
