package router

import (
	"github.com/blog-service/src/server/rest/handler"
	"github.com/blog-service/src/service"
	"github.com/gin-gonic/gin"
)

type RestBeanManager struct {
	*service.BeanFactory
}

func NewRestBeanManager(s *service.BeanFactory) *RestBeanManager {
	return &RestBeanManager{s}
}

func (rm *RestBeanManager) RegisterRoutesForAdmin(group *gin.RouterGroup) {
	adminUserHandler := handler.NewAdminuserAdminHandler(rm.GetUserService())
	adminUserRoutes := group.Group("/users")
	{
		adminUserRoutes.GET("/search/all", adminUserHandler.GetAllUsers)
		adminUserRoutes.POST("/search/all", adminUserHandler.GetAllUsersWithPagination)
		adminUserRoutes.DELETE("/:user_id", adminUserHandler.DeleteUser)
	}

}
