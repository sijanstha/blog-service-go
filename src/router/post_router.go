package router

import (
	posthandler "github.com/blog-service/src/http/post"
	postrepo "github.com/blog-service/src/repository/post"
	"github.com/blog-service/src/service/post"
	"github.com/gin-gonic/gin"
)

func (rm *RouterManager) RegisterRoutesForPost(group *gin.RouterGroup) {
	uHandler := posthandler.NewPostHandler(post.NewPostService(postrepo.NewPostRepository()))
	group.GET("/:post_id", uHandler.GetById)
	group.POST("", uHandler.Create)
	group.PUT("", uHandler.Update)
	group.POST("/search", uHandler.Get)
	group.GET("/search/all", uHandler.GetAll)
	group.POST("/search/all", uHandler.GetAllWithPagination)
	group.DELETE("/:post_id", uHandler.Delete)

}
