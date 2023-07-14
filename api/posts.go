package api

import (
	"database/sql"
	"net/http"
	"time"

	db "github.com/after23/sharing-vision-be/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createPostRequest struct {
	Title       string       `json:"title" binding:"required,min=20"`
	Content     string       `json:"content" binding:"required,min=200"`
	Category    string       `json:"category" binding:"required,min=3"`
	Status      string       `json:"status" binding:"required,oneof=publish draft thrash"`
	CreatedDate time.Time `json:"created_date"`
}

func (server *Server) createPost(ctx *gin.Context) {
	var req createPostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	req.CreatedDate  = time.Now().UTC().Add(7*time.Hour)

	arg := db.CreatePostParams{
		Title: req.Title,
		Content: req.Content,
		Category: req.Category,
		Status: req.Status,
		CreatedDate: sql.NullTime{
			Time: req.CreatedDate,
			Valid: true,
		},
	}

	_, err := server.CreatePost(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return
}