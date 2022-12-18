package main

import (
	"os"

	db "github.com/blog-service/src/datasources/mysql"
	"github.com/blog-service/src/repository"
	"github.com/blog-service/src/server/rest"
	"github.com/blog-service/src/server/rest/router"
	"github.com/blog-service/src/service"
)

func main() {
	var dbPort db.DbPort

	var roleRepo repository.IRoleRepository
	var userRepo repository.IUserRepository
	var postRepo repository.IPostRepository
	var commentRepo repository.ICommentRepository

	var roleService service.IRoleService
	var userService service.IUserService
	var userAuthService service.IUserAuthService
	var postService service.IPostService
	var commentService service.ICommentService

	dbaseDriver := os.Getenv("DB_DRIVER")
	dbaseConnectionString := os.Getenv("DB_URL")

	dbPort = db.NewDbAdapter(dbaseDriver, dbaseConnectionString)
	roleRepo = repository.NewRoleRepository(dbPort)
	userRepo = repository.NewUserRepository(dbPort)
	postRepo = repository.NewPostRepository(dbPort)
	commentRepo = repository.NewCommentRepository(dbPort)

	roleService = service.NewRoleService(roleRepo, userRepo)
	userService = service.NewUserService(userRepo, roleRepo)
	postService = service.NewPostService(postRepo)
	commentService = service.NewCommentService(postRepo, commentRepo)
	userAuthService = service.NewUserAuthService(userRepo, roleRepo)

	rm := router.NewRouterManager(roleService, userService, userAuthService, postService, commentService)
	httpServer := rest.NewHttpServerAdapter(rm)
	httpServer.StartApplication(":9090")
}
