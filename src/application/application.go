package application

import (
	commenthandler "github.com/blog-service/src/http/comment"
	posthandler "github.com/blog-service/src/http/post"
	commentrepo "github.com/blog-service/src/repository/comment"
	postrepo "github.com/blog-service/src/repository/post"
	"github.com/blog-service/src/service/comment"
	"github.com/blog-service/src/service/post"
	"github.com/gin-gonic/gin"
)

var (
	router   = gin.Default()
	v1Routes = router.Group("/api/v1")
)

func StartApplication() {

	registerRoutesForPost()
	registerRoutesForComment()
	router.Run(":8080")
}

func registerRoutesForPost() {
	pHandler := posthandler.NewPostHandler(post.NewPostService(postrepo.NewPostRepository()))

	postRoutes := v1Routes.Group("/posts")
	{
		postRoutes.GET("/:post_id", pHandler.GetById)
		postRoutes.POST("", pHandler.Create)
		postRoutes.PUT("", pHandler.Update)
		postRoutes.POST("/search", pHandler.Get)
		postRoutes.GET("/search/all", pHandler.GetAll)
		postRoutes.POST("/search/all", pHandler.GetAllWithPagination)
		postRoutes.DELETE("/:post_id", pHandler.Delete)
	}

}

func registerRoutesForComment() {
	cHandler := commenthandler.NewCommentHandler(comment.NewCommentService(postrepo.NewPostRepository(), commentrepo.NewCommentRepository()))

	commentRoutes := v1Routes.Group("/posts/:post_id/comments")
	{
		commentRoutes.GET("/:comment_id", cHandler.GetById)
		commentRoutes.POST("", cHandler.Create)
		commentRoutes.PUT("", cHandler.Update)
		commentRoutes.POST("/search", cHandler.Get)
		commentRoutes.GET("/search/all", cHandler.GetAll)
		commentRoutes.POST("/search/all", cHandler.GetAllWithPagination)
		commentRoutes.DELETE("/:comment_id", cHandler.Delete)
	}
}
