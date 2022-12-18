package service

import (
	"github.com/blog-service/src/domain/user"
	"github.com/blog-service/src/repository"
	"github.com/blog-service/src/security/jwt"
	"github.com/blog-service/src/utils"
	"github.com/blog-service/src/utils/crypto"
	"github.com/blog-service/src/utils/errors"
	stringutils "github.com/blog-service/src/utils/string"
)

type IUserAuthService interface {
	Register(*user.UserDomain) (*user.UserDomain, *errors.RestErr)
	Login(*user.UserLoginRequest) (*user.UserLoginResponse, *errors.RestErr)
}

type userAuthService struct {
	userRepo repository.IUserRepository
	roleRepo repository.IRoleRepository
}

func NewUserAuthService(userRepo repository.IUserRepository, roleRepo repository.IRoleRepository) IUserAuthService {
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

	fetchedRole, err := service.roleRepo.FindById(fetchedUser.RoleId)
	if err != nil {
		return nil, errors.NewUnauthorizedError("invalid email or password")
	}

	jwtService := &jwt.JwtTokenService{}
	token, err := jwtService.GetToken(jwt.Payload{
		Id:    fetchedUser.Id,
		Email: fetchedUser.Email,
		Role:  fetchedRole.RoleName,
	})
	if err != nil {
		return nil, errors.NewUnauthorizedError("couldn't generate token")
	}

	response := user.UserLoginResponse{
		UserDetails: *fetchedUser,
		Token:       token,
	}

	return &response, nil
}

func (service *userAuthService) findUserByEmail(email string) (*user.UserDomain, error) {
	userFilter := user.UserFilter{
		Email:   email,
		Active:  utils.BoolAddr(true),
		Deleted: utils.BoolAddr(false),
	}
	return service.userRepo.Find(userFilter)
}

func (service *userAuthService) findUserForAuthentication(email string, passwordHash string) (*user.UserDomain, error) {
	userFilter := user.UserFilter{
		Email:    email,
		Password: passwordHash,
		Active:   utils.BoolAddr(true),
		Deleted:  utils.BoolAddr(false),
	}
	return service.userRepo.Find(userFilter)
}
