package router

import (
	userhandler "github.com/blog-service/src/http/user"
	rolerepo "github.com/blog-service/src/repository/role"
	userrepo "github.com/blog-service/src/repository/user"
	"github.com/blog-service/src/service/user"
	"github.com/gin-gonic/gin"
)

func (rm *RouterManager) RegisterRoutesForUserAuth(group *gin.RouterGroup) {
	authHandler := userhandler.NewUserAuthHandler(user.NewUserAuthService(userrepo.NewUserRepository(), rolerepo.NewRoleRepository()))
	group.POST("/register", authHandler.Register)
	group.POST("/login", authHandler.Login)
}
