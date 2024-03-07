package cmd

import (
	"authen-system/internal/config"
	"authen-system/internal/controllers"
	"authen-system/internal/database"
	"authen-system/pkg/cache"
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
	voucherRepo := database.NewVoucherRepository(db)
	campaignRepo := database.NewCampaignRepository(db)
	voucherHandler := controllers.NewVoucherHandler(voucherRepo, campaignRepo)
	campaignQueue := controllers.NewCampaignQueue(voucherHandler)
	go campaignQueue.Start()

	campaignCache := cache.NewCampaign()
	handler := controllers.NewUserHandler(userRepo, config.Authentication, campaignCache, campaignQueue)
	userAPI := g.Group("/user")
	{
		userAPI.POST("/login", handler.Login)
		userAPI.POST("/signup", handler.SignUp)
	}
}
