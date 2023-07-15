package api

import (
	"time"

	db "github.com/after23/sharing-vision-be/db/sqlc"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	*db.Queries
	router *gin.Engine
}

// NewServer creates a new httpserver and setup routing.
func NewServer(q *db.Queries) *Server {
	server := &Server{Queries: q}
	router := gin.Default()

	router.Use(cors.New(cors.Config{
    AllowAllOrigins: true,
    AllowMethods:     []string{"GET", "PATCH", "POST", "DELETE", "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Content-Type"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
    MaxAge: 12 * time.Hour,
  }))
	router.POST("/article", server.createPost)
	router.GET("/article/:id", server.getPost)
	router.GET("/article/:id/:offset", server.listPost)
	router.GET("/article/:id/:offset/published", server.listPublishedPost)
	router.PATCH("/article/:id", server.updatePost)
	router.DELETE("/article/:id", server.deletePost)

	server.router = router
	return server
}

// Start run the http server on the specified address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}