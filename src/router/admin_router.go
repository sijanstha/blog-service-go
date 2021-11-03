package router

import (
	userhandler "github.com/blog-service/src/http/user"
	rolerepo "github.com/blog-service/src/repository/role"
	userrepo "github.com/blog-service/src/repository/user"
	"github.com/blog-service/src/service/user"
	"github.com/gin-gonic/gin"
)

func (rm *RouterManager) RegisterRoutesForAdmin(group *gin.RouterGroup) {
	adminUserHandler := userhandler.NewAdminuserAdminHandler(user.NewUserService(userrepo.NewUserRepository(), rolerepo.NewRoleRepository()))
	adminUserRoutes := group.Group("/users")
	{
		adminUserRoutes.GET("/search/all", adminUserHandler.GetAllUsers)
		adminUserRoutes.POST("/search/all", adminUserHandler.GetAllUsersWithPagination)
		adminUserRoutes.DELETE("/:user_id", adminUserHandler.DeleteUser)
	}

}
