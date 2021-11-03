package application

import (
	routermanager "github.com/blog-service/src/router"
	"github.com/blog-service/src/security/middleware"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {

	router.Use(middleware.CORSMiddleware())
	rootGroup := router.Group("/api/v1")
	mm := routermanager.RouterManager{}
	mm.RegisterRoutesForUserAuth(rootGroup.Group(""))
	mm.RegisterRoutesForRole(rootGroup.Group("/roles"))
	mm.RegisterRoutesForPost(rootGroup.Group("/posts"))
	mm.RegisterRoutesForComment(rootGroup.Group("/posts/:post_id/comments"))
	mm.RegisterRoutesForUser(rootGroup.Group("/users"))
	mm.RegisterRoutesForAdmin(rootGroup.Group("/admin"))

	router.Run(":8080")
}
