package router

import (
	commenthandler "github.com/blog-service/src/http/comment"
	commentrepo "github.com/blog-service/src/repository/comment"
	postrepo "github.com/blog-service/src/repository/post"
	"github.com/blog-service/src/service/comment"
	"github.com/gin-gonic/gin"
)

func (rm *RouterManager) RegisterRoutesForComment(group *gin.RouterGroup) {
	cHandler := commenthandler.NewCommentHandler(comment.NewCommentService(postrepo.NewPostRepository(), commentrepo.NewCommentRepository()))
	group.GET("/:comment_id", cHandler.GetById)
	group.POST("", cHandler.Create)
	group.PUT("", cHandler.Update)
	group.POST("/search", cHandler.Get)
	group.GET("/search/all", cHandler.GetAll)
	group.POST("/search/all", cHandler.GetAllWithPagination)
	group.DELETE("/:comment_id", cHandler.Delete)

}
