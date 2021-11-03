package application

import (
	commenthandler "github.com/blog-service/src/http/comment"
	posthandler "github.com/blog-service/src/http/post"
	rolehandler "github.com/blog-service/src/http/role"
	userhandler "github.com/blog-service/src/http/user"
	commentrepo "github.com/blog-service/src/repository/comment"
	postrepo "github.com/blog-service/src/repository/post"
	rolerepo "github.com/blog-service/src/repository/role"
	userrepo "github.com/blog-service/src/repository/user"
	"github.com/blog-service/src/security/middleware"
	"github.com/blog-service/src/service/comment"
	"github.com/blog-service/src/service/post"
	"github.com/blog-service/src/service/role"
	"github.com/blog-service/src/service/user"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
	// v1Routes = router.Group("/api/v1")
)

func StartApplication() {

	router.Use(middleware.CORSMiddleware())
	registerRoutesForUserAuth()
	registerRoutesForRole()
	registerRoutesForPost()
	registerRoutesForComment()
	registerRoutesForUser()
	router.Run(":8080")
}

func registerRoutesForPost() {
	uHandler := posthandler.NewPostHandler(post.NewPostService(postrepo.NewPostRepository()))

	postRoutes := router.Group("/api/v1/posts")
	{
		postRoutes.GET("/:post_id", uHandler.GetById)
		postRoutes.POST("", uHandler.Create)
		postRoutes.PUT("", uHandler.Update)
		postRoutes.POST("/search", uHandler.Get)
		postRoutes.GET("/search/all", uHandler.GetAll)
		postRoutes.POST("/search/all", uHandler.GetAllWithPagination)
		postRoutes.DELETE("/:post_id", uHandler.Delete)
	}
}

func registerRoutesForComment() {
	cHandler := commenthandler.NewCommentHandler(comment.NewCommentService(postrepo.NewPostRepository(), commentrepo.NewCommentRepository()))

	commentRoutes := router.Group("/api/v1/posts/:post_id/comments")
	{
		commentRoutes.GET("/:comment_id", cHandler.GetById)
		commentRoutes.POST("", cHandler.Create)
		commentRoutes.PUT("", cHandler.Update)
		commentRoutes.POST("/search", cHandler.Get)
		commentRoutes.GET("/search/all", cHandler.GetAll)
		commentRoutes.POST("/search/all", cHandler.GetAllWithPagination)
		commentRoutes.DELETE("/:comment_id", cHandler.Delete)
	}
}

func registerRoutesForRole() {
	rHandler := rolehandler.NewRoleHandler(role.NewRoleService(rolerepo.NewRoleRepository(), userrepo.NewUserRepository()))

	roleRoutes := router.Group("/api/v1/roles")
	{
		roleRoutes.GET("/:role_id", rHandler.GetById)
		roleRoutes.POST("", rHandler.Create)
		roleRoutes.PUT("", rHandler.Update)
		roleRoutes.POST("/search", rHandler.Get)
		roleRoutes.GET("/search/all", rHandler.GetAll)
		roleRoutes.DELETE("/:role_id", rHandler.Delete)
	}
}

func registerRoutesForUser() {
	uHandler := userhandler.NewUserHandler(user.NewUserService(userrepo.NewUserRepository(), rolerepo.NewRoleRepository()))

	postRoutes := router.Group("/api/v1/users")
	{
		postRoutes.GET("/:user_id", uHandler.GetById)
		postRoutes.PUT("/profile", uHandler.Update)
		postRoutes.POST("/search", uHandler.Get)
	}

	adminUserHandler := userhandler.NewAdminuserAdminHandler(user.NewUserService(userrepo.NewUserRepository(), rolerepo.NewRoleRepository()))
	adminUserRoutes := router.Group("/admin/users")
	{
		adminUserRoutes.GET("/search/all", adminUserHandler.GetAllUsers)
		adminUserRoutes.POST("/search/all", adminUserHandler.GetAllUsersWithPagination)
		adminUserRoutes.DELETE("/:user_id", adminUserHandler.DeleteUser)
	}
}

func registerRoutesForUserAuth() {
	authHandler := userhandler.NewUserAuthHandler(user.NewUserAuthService(userrepo.NewUserRepository(), rolerepo.NewRoleRepository()))
	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)
}
