package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createPostRequest struct {
	Title       string       `json:"title" binding:"required,len=gte 20"`
	Content     string       `json:"content" binding:"required,len=gte 200"`
	Category    string       `json:"category" binding:"required,len=gte 3"`
	Status      string       `json:"status" binding:"required,oneof=publish draft thrash"`
	CreatedDate sql.NullTime `json:"created_date"`
}

func (server *Server) createPost(ctx *gin.Context) {
	var req createPostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
}