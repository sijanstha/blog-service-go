package router

import (
	"github.com/blog-service/src/server/rest/handler"
	"github.com/gin-gonic/gin"
)

func (rm *RestBeanManager) RegisterRoutesForUserAuth(group *gin.RouterGroup) {
	authHandler := handler.NewUserAuthHandler(rm.GetUserAuthService())
	group.POST("/register", authHandler.Register)
	group.POST("/login", authHandler.Login)
}
