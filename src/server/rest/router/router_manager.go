package router

import "github.com/blog-service/src/service"

type RouterManager struct {
	roleService     service.IRoleService
	userService     service.IUserService
	userAuthService service.IUserAuthService
	postService     service.IPostService
	commentService  service.ICommentService
}

func NewRouterManager(roleService service.IRoleService,
	userService service.IUserService,
	userAuthService service.IUserAuthService,
	postService service.IPostService,
	commentService service.ICommentService) *RouterManager {
	return &RouterManager{roleService, userService, userAuthService, postService, commentService}
}

func (rm RouterManager) GetRoleService() service.IRoleService {
	return rm.roleService
}

func (rm RouterManager) GetUserService() service.IUserService {
	return rm.userService
}

func (rm RouterManager) GetPostService() service.IPostService {
	return rm.postService
}
