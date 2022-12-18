package router

import (
	"github.com/blog-service/src/server/rest/handler"
	"github.com/gin-gonic/gin"
)

func (rm *RouterManager) RegisterRoutesForUserAuth(group *gin.RouterGroup) {
	authHandler := handler.NewUserAuthHandler(rm.userAuthService)
	group.POST("/register", authHandler.Register)
	group.POST("/login", authHandler.Login)
}
