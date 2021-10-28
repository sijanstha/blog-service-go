package application

import (
	posthandler "github.com/blog-service/src/http/post"
	postrepo "github.com/blog-service/src/repository/post"
	"github.com/blog-service/src/service/post"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {

	registerRoutesForPost()
	router.Run(":8080")
}

func registerRoutesForPost() {
	pHandler := posthandler.NewPostHandler(post.NewPostService(postrepo.NewPostRepository()))

	router.GET("/posts/:post_id", pHandler.GetById)
	router.POST("/posts", pHandler.Create)
	router.PUT("/posts", pHandler.Update)
	router.POST("/posts/search", pHandler.Get)
	router.GET("/posts/search/all", pHandler.GetAll)
	router.POST("/posts/search/all", pHandler.GetAllWithPagination)
}
