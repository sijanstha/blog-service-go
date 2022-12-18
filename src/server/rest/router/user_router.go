package router

import (
	"github.com/blog-service/src/server/rest/handler"
	"github.com/gin-gonic/gin"
)

func (rm *RouterManager) RegisterRoutesForUser(group *gin.RouterGroup) {
	uHandler := handler.NewUserHandler(rm.userService)

	group.GET("/:user_id", uHandler.GetById)
	group.PUT("/profile", uHandler.Update)
	group.POST("/search", uHandler.Get)

}
