package router

import (
	"github.com/blog-service/src/server/rest/handler"
	"github.com/gin-gonic/gin"
)

func (rm *RouterManager) RegisterRoutesForComment(group *gin.RouterGroup) {
	cHandler := handler.NewCommentHandler(rm.commentService)
	group.GET("/:comment_id", cHandler.GetById)
	group.POST("", cHandler.Create)
	group.PUT("", cHandler.Update)
	group.POST("/search", cHandler.Get)
	group.GET("/search/all", cHandler.GetAll)
	group.POST("/search/all", cHandler.GetAllWithPagination)
	group.DELETE("/:comment_id", cHandler.Delete)

}
