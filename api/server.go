package api

import (
	db "github.com/after23/sharing-vision-be/db/sqlc"
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

	router.POST("/article", server.createPost)
	router.GET("/article/:id", server.getPost)
	router.GET("/article/:id/:offset", server.listPost)

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