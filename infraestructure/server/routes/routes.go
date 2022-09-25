package routes

import (
	"github.com/gin-gonic/gin"
)

// @title           Todo List Golang
// @version         1.0
// @description     Todo List Golang swagger server
// @termsOfService  http://swagger.io/terms/

// @host      api-hml.orbitpages.com/workspaces
// @BasePath  /v1
func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("v1") //TODO logger
	{
		item := main.Group("item")
		{
			item.POST("/")
		}
	}

	return router
}
