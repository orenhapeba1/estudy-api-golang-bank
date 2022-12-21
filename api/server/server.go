package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/orenhapeba1/estudy-api-golang-bank/routes"
	"log"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")
	router.Use(cors.New(config))

	routes.ConfigRoutes(router)

	return Server{
		port:   "5000",
		server: router,
	}
}

func (s *Server) Run() {
	log.Printf("Server running at port: %v", s.port)
	log.Fatal(s.server.Run(":" + s.port))
}
