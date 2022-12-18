package service

type BeanFactory struct {
	roleService     IRoleService
	userService     IUserService
	userAuthService IUserAuthService
	postService     IPostService
	commentService  ICommentService
}

func CreateFactory(roleService IRoleService,
	userService IUserService,
	userAuthService IUserAuthService,
	postService IPostService,
	commentService ICommentService) *BeanFactory {
	return &BeanFactory{roleService, userService, userAuthService, postService, commentService}
}

func (bf BeanFactory) GetRoleService() IRoleService {
	return bf.roleService
}

func (bf BeanFactory) GetUserService() IUserService {
	return bf.userService
}

func (bf BeanFactory) GetPostService() IPostService {
	return bf.postService
}

func (bf BeanFactory) GetUserAuthService() IUserAuthService {
	return bf.userAuthService
}

func (bf BeanFactory) GetCommentService() ICommentService {
	return bf.commentService
}
