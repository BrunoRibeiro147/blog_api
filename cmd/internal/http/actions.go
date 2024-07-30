package http

import (
	"net/http"

	"blogapi/cmd/internal"
	"blogapi/cmd/internal/post"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

var service post.Service

func Configure(conn *pgxpool.Pool) {
	service = post.Service{
		Repository: post.Repository{
			Conn: conn,
		},
	}
}

func PostPosts(ctx *gin.Context) {
	var post internal.Post
	if err := ctx.BindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := service.Create(post); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusCreated, nil)
}

func DeletePosts(ctx *gin.Context) {
	params := ctx.Param("id")

	if err := service.Delete(params); err != nil {
		statusCode := http.StatusInternalServerError
		if err == post.ErrPostNotFound {
			statusCode = http.StatusNotFound
		}

		ctx.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
	}
}

func GetPosts(ctx *gin.Context) {
	params := ctx.Param("id")

	p, err := service.FindOneById(params)
	if err != nil {
		statusCode := http.StatusInternalServerError

		if err == post.ErrPostNotFound {
			statusCode = http.StatusNotFound
		}

		ctx.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, p)
}

func HelloWorld(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello World",
	})
}
