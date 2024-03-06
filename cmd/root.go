package cmd

import (
	"authen-system/internal/config"
	"github.com/gin-gonic/gin"
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
}
