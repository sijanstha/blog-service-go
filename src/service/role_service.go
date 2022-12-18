package service

import (
	"github.com/blog-service/src/domain/role"
	"github.com/blog-service/src/domain/user"
	"github.com/blog-service/src/repository"
	"github.com/blog-service/src/utils"
	dateutils "github.com/blog-service/src/utils/date"
	"github.com/blog-service/src/utils/errors"
	"github.com/blog-service/src/utils/logger"
	stringutils "github.com/blog-service/src/utils/string"
)

type IRoleService interface {
	Save(*role.Role) (*role.Role, *errors.RestErr)
	Update(*role.Role) (*role.Role, *errors.RestErr)
	FindById(string) (*role.Role, *errors.RestErr)
	Find(role.RoleFilter) (*role.Role, *errors.RestErr)
	FindAll() []role.Role
	Delete(string) *errors.RestErr
}

type roleService struct {
	roleRepository repository.IRoleRepository
	userRepository repository.IUserRepository
}

func NewRoleService(roleRepository repository.IRoleRepository, userRepository repository.IUserRepository) IRoleService {
	return &roleService{
		roleRepository,
		userRepository,
	}
}

func (s *roleService) Save(request *role.Role) (*role.Role, *errors.RestErr) {
	err := request.Validate()
	if err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}

	request.Id = stringutils.GenerateUniqueId()
	request.CreatedAt = dateutils.GetNow().String()
	request.UpdatedAt = request.CreatedAt
	request.DeletedAt = ""
	request.IsActive = true
	request.IsDeleted = false
	if stringutils.IsEmptyOrNull(request.DisplayName) {
		request.DisplayName = request.RoleName
	}

	roleFilter := role.RoleFilter{
		RoleName: request.RoleName,
		Active:   func(b bool) *bool { return &b }(true),
		Deleted:  func(b bool) *bool { return &b }(false),
	}
	fetchedRole, _ := s.roleRepository.Find(roleFilter)
	if fetchedRole != nil {
		return nil, errors.NewBadRequestError("role name already exists")
	}

	result, err := s.roleRepository.Save(request)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return result, nil
}

func (s *roleService) Update(request *role.Role) (*role.Role, *errors.RestErr) {
	if request.Id == "" || len(request.Id) <= 0 {
		return nil, errors.NewBadRequestError("role id missing")
	}

	err := request.Validate()
	if err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}

	_, restErr := s.FindById(request.Id)
	if restErr != nil {
		return nil, restErr
	}

	request.UpdatedAt = dateutils.GetNow().String()
	request.IsActive = true
	request.IsDeleted = false

	result, err := s.roleRepository.Update(request)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return result, nil
}

func (s *roleService) FindById(id string) (*role.Role, *errors.RestErr) {
	result, err := s.roleRepository.FindById(id)
	if err != nil {
		if err == repository.ErrRoleNotFound {
			return nil, errors.NewNotFoundError(err.Error())
		}
		return nil, errors.NewInternalServerError(err.Error())
	}

	return result, nil
}

func (s *roleService) Find(filter role.RoleFilter) (*role.Role, *errors.RestErr) {
	if err := filter.Validate(); err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}
	result, err := s.roleRepository.Find(filter)
	if err != nil {
		if err == repository.ErrRoleNotFound {
			return nil, errors.NewNotFoundError(err.Error())
		}
		return nil, errors.NewInternalServerError(err.Error())
	}

	return result, nil
}

func (s *roleService) FindAll() []role.Role {
	return s.roleRepository.FindAll()
}

func (s *roleService) Delete(id string) *errors.RestErr {
	req, restErr := s.FindById(id)
	if restErr != nil {
		return restErr
	}

	count, err := s.userRepository.CountUser(user.UserFilter{
		RoleId:  id,
		Active:  utils.BoolAddr(true),
		Deleted: utils.BoolAddr(false),
	})
	if err != nil {
		if err != repository.ErrUserNotFound {
			logger.Error(err.Error(), err)
			return errors.NewInternalServerError("server down")
		}
	}
	if count >= 1 {
		return errors.NewBadRequestError("cannot delete this role. role already in use")
	}

	err = s.roleRepository.Delete(req)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
