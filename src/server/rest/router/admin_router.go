package router

import (
	"github.com/blog-service/src/server/rest/handler"
	"github.com/gin-gonic/gin"
)

func (rm *RouterManager) RegisterRoutesForAdmin(group *gin.RouterGroup) {
	adminUserHandler := handler.NewAdminuserAdminHandler(rm.userService)
	adminUserRoutes := group.Group("/users")
	{
		adminUserRoutes.GET("/search/all", adminUserHandler.GetAllUsers)
		adminUserRoutes.POST("/search/all", adminUserHandler.GetAllUsersWithPagination)
		adminUserRoutes.DELETE("/:user_id", adminUserHandler.DeleteUser)
	}

}
