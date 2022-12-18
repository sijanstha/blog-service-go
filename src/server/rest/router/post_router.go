package router

import (
	"github.com/blog-service/src/server/rest/handler"
	"github.com/gin-gonic/gin"
)

func (rm *RestBeanManager) RegisterRoutesForPost(group *gin.RouterGroup) {
	uHandler := handler.NewPostHandler(rm.GetPostService())
	group.GET("/:post_id", uHandler.GetById)
	group.POST("", uHandler.Create)
	group.PUT("", uHandler.Update)
	group.POST("/search", uHandler.Get)
	group.GET("/search/all", uHandler.GetAll)
	group.POST("/search/all", uHandler.GetAllWithPagination)
	group.DELETE("/:post_id", uHandler.Delete)

}
