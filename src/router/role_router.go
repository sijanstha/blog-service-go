package router

import (
	rolehandler "github.com/blog-service/src/http/role"
	rolerepo "github.com/blog-service/src/repository/role"
	userrepo "github.com/blog-service/src/repository/user"
	"github.com/blog-service/src/service/role"
	"github.com/gin-gonic/gin"
)

func (rm *RouterManager) RegisterRoutesForRole(group *gin.RouterGroup) {
	rHandler := rolehandler.NewRoleHandler(role.NewRoleService(rolerepo.NewRoleRepository(), userrepo.NewUserRepository()))

	group.GET("/:role_id", rHandler.GetById)
	group.POST("", rHandler.Create)
	group.PUT("", rHandler.Update)
	group.POST("/search", rHandler.Get)
	group.GET("/search/all", rHandler.GetAll)
	group.DELETE("/:role_id", rHandler.Delete)

}
