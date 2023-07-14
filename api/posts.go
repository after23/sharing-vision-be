package api

import (
	"database/sql"
	"net/http"
	"time"

	db "github.com/after23/sharing-vision-be/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createPostRequest struct {
	Title       string `json:"title" binding:"required,min=20"`
	Content     string `json:"content" binding:"required,min=200"`
	Category    string `json:"category" binding:"required,min=3"`
	Status      string `json:"status" binding:"required,oneof=publish draft thrash"`
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

type getPostRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getPost(ctx *gin.Context) {
	var req getPostRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	post, err := server.GetPostById(ctx, req.ID)
	
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, post)
	return
}

type listPostRequest struct {
	PageId int32 `uri:"offset" binding:"required,min=1"`
	PageSize int32 `uri:"id" binding:"required,min=5"`
}

func (server *Server) listPost(ctx *gin.Context) {
	var req listPostRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.GetPostParams{
		Limit: req.PageSize,
		Offset: (req.PageId-1)*req.PageSize,
	}

	posts, err := server.GetPost(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, posts)
	return
}

type updatePostRequest struct {
	Title       string `json:"title" binding:"required,min=20"`
	Content     string `json:"content" binding:"required,min=200"`
	Category    string `json:"category" binding:"required,min=3"`
	Status      string `json:"status" binding:"required,oneof=publish draft thrash"`
	UpdatedDate time.Time `json:"created_date"`
}

func (server *Server) updatePost(ctx *gin.Context) {
	var req updatePostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	var post getPostRequest	
	if err := ctx.ShouldBindUri(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	req.UpdatedDate = time.Now().UTC().Add(7*time.Hour)
	

	arg := db.UpdatePostParams{
		Title: req.Title,
		Content: req.Content,
		Status: req.Status,
		Category: req.Category,
		ID: post.ID,
		UpdatedDate: sql.NullTime{
			Time: req.UpdatedDate,
			Valid: true,
		},
	}

	if _, err := server.GetPostById(ctx, post.ID); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}
	_, err := server.UpdatePost(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
	return
}