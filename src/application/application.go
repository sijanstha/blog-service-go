package application

import (
	posthandler "github.com/blog-service/src/http/post"
	postrepo "github.com/blog-service/src/repository/post"
	"github.com/blog-service/src/service/post"
	"github.com/gin-gonic/gin"
)

var (
	router   = gin.Default()
	v1Routes = router.Group("/api/v1")
)

func StartApplication() {

	registerRoutesForPost()
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
	}

}
