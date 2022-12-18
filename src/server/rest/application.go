package rest

import (
	"github.com/blog-service/src/security/middleware"
	"github.com/blog-service/src/server"
	"github.com/blog-service/src/server/rest/router"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	engine *gin.Engine
}

func NewHttpServerAdapter(rm *router.RestBeanManager) server.ServerPort {
	var engine *gin.Engine = gin.Default()
	engine.Use(middleware.CORSMiddleware())
	rootGroup := engine.Group("/api/v1")

	rm.RegisterRoutesForUserAuth(rootGroup.Group(""))
	rm.RegisterRoutesForRole(rootGroup.Group("/roles"))
	rm.RegisterRoutesForPost(rootGroup.Group("/posts"))
	rm.RegisterRoutesForComment(rootGroup.Group("/posts/:post_id/comments"))
	rm.RegisterRoutesForUser(rootGroup.Group("/users"))
	rm.RegisterRoutesForAdmin(rootGroup.Group("/admin"))

	return &HttpServer{engine: engine}
}

func (server HttpServer) StartApplication(addr string) {
	server.engine.Run(addr)
}
