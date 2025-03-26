package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maretrodep/base-auth-go/jwt/config"
	"github.com/maretrodep/base-auth-go/jwt/internal/handlers"
	"gorm.io/gorm"
)

type Server struct {
	router    *gin.Engine
	db        *gorm.DB
	serverCfg *config.ServerConfig
	authCfg   *config.AuthConfig
}

func NewServer(db *gorm.DB, serverCfg *config.ServerConfig, authCfg *config.AuthConfig) *Server {
	router := gin.Default()

	authHandler := handlers.NewAuthHandler(db, authCfg.JWTSecret)

	public := router.Group("/api")
	public.POST("/signup", authHandler.Signup)
	public.POST("/hello", func(c *gin.Context) {
		println("Hello world")
		c.JSON(http.StatusOK, gin.H{})
	})

	return &Server{
		router:    router,
		db:        db,
		serverCfg: serverCfg,
		authCfg:   authCfg,
	}
}

// Start runs the Gin server.
func (s *Server) Start() error {
	addr := fmt.Sprintf(":%s", s.serverCfg.Port)
	return s.router.Run(addr)
}
