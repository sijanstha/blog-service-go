package router

import (
	"github.com/blog-service/src/server/rest/handler"
	"github.com/gin-gonic/gin"
)

func (rm *RestBeanManager) RegisterRoutesForRole(group *gin.RouterGroup) {
	rHandler := handler.NewRoleHandler(rm.GetRoleService())

	group.GET("/:role_id", rHandler.GetById)
	group.POST("", rHandler.Create)
	group.PUT("", rHandler.Update)
	group.POST("/search", rHandler.Get)
	group.GET("/search/all", rHandler.GetAll)
	group.DELETE("/:role_id", rHandler.Delete)

}
