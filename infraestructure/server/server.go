package server

import (
	"log"
	"net/http"

	"github.com/luisfernandomoraes/todo-list-golang/infraestructure/server/routes"

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

	router.GET("/hc", func(c *gin.Context) {
		c.String(http.StatusOK, "healthy")
	})

	log.Fatal(router.Run(":" + s.port))
}
