package role

import (
	"github.com/blog-service/src/domain/role"
	rolerepo "github.com/blog-service/src/repository/role"
	dateutils "github.com/blog-service/src/utils/date"
	"github.com/blog-service/src/utils/errors"
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
	roleRepository rolerepo.IRoleRepository
}

func NewRoleService(roleRepository rolerepo.IRoleRepository) IRoleService {
	return &roleService{
		roleRepository,
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
		if err == rolerepo.ErrRoleNotFound {
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
		if err == rolerepo.ErrRoleNotFound {
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

	err := s.roleRepository.Delete(req)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
