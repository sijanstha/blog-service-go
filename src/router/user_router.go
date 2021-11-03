package router

import (
	userhandler "github.com/blog-service/src/http/user"
	rolerepo "github.com/blog-service/src/repository/role"
	userrepo "github.com/blog-service/src/repository/user"
	"github.com/blog-service/src/service/user"
	"github.com/gin-gonic/gin"
)

func (rm *RouterManager) RegisterRoutesForUser(group *gin.RouterGroup) {
	uHandler := userhandler.NewUserHandler(user.NewUserService(userrepo.NewUserRepository(), rolerepo.NewRoleRepository()))

	group.GET("/:user_id", uHandler.GetById)
	group.PUT("/profile", uHandler.Update)
	group.POST("/search", uHandler.Get)

}
