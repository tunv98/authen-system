package cmd

import (
	"authen-system/internal/auth"
	"authen-system/internal/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateGin() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/health"},
	}), gin.Recovery())
	return engine
}

func RegisterHandler(
	r gin.IRouter,
	config config.App,
) {
	apiV1 := r.Group("/api/v1")
	registerUserHandler(apiV1, config)
	protectedAPIV1 := apiV1.Group("/home")
	protectedAPIV1.Use(auth.ValidateToken(config.Authentication.SecretKey))
	{
		protectedAPIV1.GET("", accessHomePage)
	}
}

func accessHomePage(c *gin.Context) {
	userId, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "user information is not existed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": fmt.Sprintf("you're logged in with id is %s", userId),
	})
}
