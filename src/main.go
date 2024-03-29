package main

import (
	"os"

	"github.com/blog-service/src/config"
	db "github.com/blog-service/src/datasources/mysql"
	"github.com/blog-service/src/repository"
	"github.com/blog-service/src/server"
	"github.com/blog-service/src/server/grpc"
	"github.com/blog-service/src/server/grpc/handler"
	"github.com/blog-service/src/server/rest"
	"github.com/blog-service/src/server/rest/router"
	"github.com/blog-service/src/service"
	stringutils "github.com/blog-service/src/utils/string"
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

	var server server.ServerPort
	var kafkaProducer *config.KafkaProducer

	dbaseDriver := os.Getenv("DB_DRIVER")
	dbaseConnectionString := os.Getenv("DB_URL")
	grpcFlag := os.Getenv("GRPC_FLAG")
	kafkaBootstrapServer := os.Getenv("KAFKA_BOOTSTRAP_SERVER")
	postChangeTopic := os.Getenv("POST_CHANGE_TOPIC")
	isGrpcServerFlagEnabled := (!stringutils.IsEmptyOrNull(grpcFlag) && (grpcFlag == "True" || grpcFlag == "true"))

	dbPort = db.NewDbAdapter(dbaseDriver, dbaseConnectionString)
	roleRepo = repository.NewRoleRepository(dbPort)
	userRepo = repository.NewUserRepository(dbPort)
	postRepo = repository.NewPostRepository(dbPort)
	commentRepo = repository.NewCommentRepository(dbPort)
	kafkaProducer = config.NewKafkaProducer(kafkaBootstrapServer)

	defer kafkaProducer.CloseConnection()

	roleService = service.NewRoleService(roleRepo, userRepo)
	userService = service.NewUserService(userRepo, roleRepo)
	postService = service.NewPostService(postRepo, kafkaProducer, postChangeTopic)
	commentService = service.NewCommentService(postRepo, commentRepo)
	userAuthService = service.NewUserAuthService(userRepo, roleRepo)

	rm := service.CreateFactory(roleService, userService, userAuthService, postService, commentService)
	if isGrpcServerFlagEnabled {
		server = grpc.NewGrpcServerAdapter(handler.NewGrpcBeanManager(rm))
	} else {
		server = rest.NewHttpServerAdapter(router.NewRestBeanManager(rm))
	}

	server.StartApplication(":9090")
}
