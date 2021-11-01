package user

import (
	"github.com/blog-service/src/domain/user"
	rolerepo "github.com/blog-service/src/repository/role"
	userrepo "github.com/blog-service/src/repository/user"
	"github.com/blog-service/src/utils/crypto"
	"github.com/blog-service/src/utils/errors"
	stringutils "github.com/blog-service/src/utils/string"
)

type IUserAuthService interface {
	Register(*user.UserDomain) (*user.UserDomain, *errors.RestErr)
	Login(*user.UserLoginRequest) (*user.UserLoginResponse, *errors.RestErr)
}

type userAuthService struct {
	userRepo userrepo.IUserRepository
	roleRepo rolerepo.IRoleRepository
}

func NewUserAuthService(userRepo userrepo.IUserRepository, roleRepo rolerepo.IRoleRepository) IUserAuthService {
	return &userAuthService{userRepo, roleRepo}
}

func (service *userAuthService) Register(request *user.UserDomain) (*user.UserDomain, *errors.RestErr) {
	err := request.Validate(user.SAVE_REQUEST_TYPE)
	if err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}

	_, err = service.roleRepo.FindById(request.RoleId)
	if err != nil {
		return nil, errors.NewNotFoundError(err.Error())
	}

	fetchedUser, _ := service.findUserByEmail(request.Email)
	if fetchedUser != nil {
		return nil, errors.NewBadRequestError("email already exists")
	}

	request.Id = stringutils.GenerateUniqueId()
	request.Password = crypto.GetMd5(request.Password)
	entity := request.ToUserEntity(user.SAVE_REQUEST_TYPE)

	result, err := service.userRepo.Save(&entity)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return result, nil
}

func (service *userAuthService) Login(request *user.UserLoginRequest) (*user.UserLoginResponse, *errors.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}

	fetchedUser, err := service.findUserForAuthentication(request.Email, crypto.GetMd5(request.Password))
	if err != nil {
		return nil, errors.NewUnauthorizedError("invalid email or password")
	}

	response := user.UserLoginResponse{
		UserDetails: *fetchedUser,
		Token:       "ABCDEFGHIJKLMNOPQRST",
	}

	return &response, nil
}

func (service *userAuthService) findUserByEmail(email string) (*user.UserDomain, error) {
	userFilter := user.UserFilter{
		Email:   email,
		Active:  func(b bool) *bool { return &b }(true),
		Deleted: func(b bool) *bool { return &b }(false),
	}
	return service.userRepo.Find(userFilter)
}

func (service *userAuthService) findUserForAuthentication(email string, passwordHash string) (*user.UserDomain, error) {
	userFilter := user.UserFilter{
		Email:    email,
		Password: passwordHash,
		Active:   func(b bool) *bool { return &b }(true),
		Deleted:  func(b bool) *bool { return &b }(false),
	}
	return service.userRepo.Find(userFilter)
}
