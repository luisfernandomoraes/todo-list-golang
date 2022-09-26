package server

import (
	"github.com/luisfernandomoraes/todo-list-golang/infraestructure/server/middlewares"
	"github.com/luisfernandomoraes/todo-list-golang/infraestructure/server/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "5100",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	router.Use(middlewares.JSONLogMiddleware())

	router.GET("/hc", func(c *gin.Context) {
		c.String(http.StatusOK, "healthy")
	})

	log.Fatal(router.Run(":" + s.port))
}
