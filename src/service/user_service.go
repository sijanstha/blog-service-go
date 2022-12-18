package service

import (
	"github.com/blog-service/src/domain/user"
	"github.com/blog-service/src/repository"
	"github.com/blog-service/src/utils/errors"
)

type IUserService interface {
	Update(*user.UserDomain) (*user.UserDomain, *errors.RestErr)
	FindById(string) (*user.UserDomain, *errors.RestErr)
	Find(user.UserFilter) (*user.UserDomain, *errors.RestErr)
	FindAll() user.UserDomainList
	FindAllWithPagination(user.UserListFilter) (*user.UserPaginationDetails, *errors.RestErr)
	Delete(string) *errors.RestErr
}

type userService struct {
	userRepo repository.IUserRepository
	roleRepo repository.IRoleRepository
}

func NewUserService(userRepo repository.IUserRepository, roleRepo repository.IRoleRepository) IUserService {
	return &userService{
		userRepo,
		roleRepo,
	}
}

func (s *userService) Update(request *user.UserDomain) (*user.UserDomain, *errors.RestErr) {
	if request.Id == "" || len(request.Id) <= 0 {
		return nil, errors.NewBadRequestError("user id missing")
	}

	err := request.Validate(user.UPDATE_REQUEST_TYPE)
	if err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}

	_, restErr := s.FindById(request.Id)
	if restErr != nil {
		return nil, restErr
	}

	_, err = s.roleRepo.FindById(request.RoleId)
	if err != nil {
		return nil, errors.NewNotFoundError(err.Error())
	}

	entity := request.ToUserEntity(user.UPDATE_REQUEST_TYPE)

	result, err := s.userRepo.Update(&entity)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return result, nil
}

func (s *userService) FindById(id string) (*user.UserDomain, *errors.RestErr) {
	result, err := s.userRepo.FindById(id)
	if err != nil {
		if err == repository.ErrUserNotFound {
			return nil, errors.NewNotFoundError(err.Error())
		}
		return nil, errors.NewInternalServerError(err.Error())
	}

	return result, nil
}

func (s *userService) Find(filter user.UserFilter) (*user.UserDomain, *errors.RestErr) {
	if err := filter.Validate(); err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}
	result, err := s.userRepo.Find(filter)
	if err != nil {
		if err == repository.ErrUserNotFound {
			return nil, errors.NewNotFoundError(err.Error())
		}
		return nil, errors.NewInternalServerError(err.Error())
	}

	return result, nil
}

func (s *userService) FindAll() user.UserDomainList {
	return s.userRepo.FindAll()
}

func (s *userService) FindAllWithPagination(filter user.UserListFilter) (*user.UserPaginationDetails, *errors.RestErr) {
	if err := filter.Validate(); err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}
	return s.userRepo.FindAllWithPagination(filter), nil
}

func (s *userService) Delete(id string) *errors.RestErr {
	req, restErr := s.FindById(id)
	if restErr != nil {
		return restErr
	}
	entity := req.ToUserEntity(user.UPDATE_REQUEST_TYPE)
	err := s.userRepo.Delete(&entity)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
