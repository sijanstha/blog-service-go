package user

import (
	"errors"

	dateutils "github.com/blog-service/src/utils/date"
	stringutils "github.com/blog-service/src/utils/string"
)

const (
	SAVE_REQUEST_TYPE = iota
	UPDATE_REQUEST_TYPE
)

var (
	ErrMissingFirstName = errors.New("first name required")
	ErrMissingLastName  = errors.New("last name required")
	ErrMissingEmail     = errors.New("email required")
	ErrMissingPassword  = errors.New("password required")
	ErrMissingRole      = errors.New("role required")
	ErrInvalidFirstName = errors.New("invalid first name")
	ErrInvalidLastName  = errors.New("invalid last name")
	ErrInvalidEmail     = errors.New("invalid email")
	ErrInvalidPassword  = errors.New("invalid password")
)

type User struct {
	Id                string `json:"id"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	Email             string `json:"email"`
	PasswordHash      string `json:"passwordHash"`
	ProfilePictureUrl string `json:"profilePictureUrl"`
	RoleId            string `json:"roleId"`
	IsActive          bool   `json:"active"`
	IsDeleted         bool   `json:"deleted"`
	CreatedAt         string `json:"createdAt"`
	UpdatedAt         string `json:"updatedAt"`
	DeletedAt         string `json:"deletedAt"`
}

type UserDomain struct {
	Id                string `json:"id"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	ProfilePictureUrl string `json:"profilePictureUrl"`
	RoleId            string `json:"roleId"`
	IsActive          bool   `json:"active"`
	IsDeleted         bool   `json:"deleted"`
	CreatedAt         string `json:"createdAt"`
	UpdatedAt         string `json:"updatedAt"`
	DeletedAt         string `json:"deletedAt"`
}

type UserPaginationDetails struct {
	Size  int            `json:"pageSize"`
	Page  int            `json:"page"`
	Total int            `json:"totalRecords"`
	Data  UserDomainList `json:"data"`
}

type UserList []User
type UserDomainList []UserDomain

func (u *UserDomain) Validate(requestType int) error {
	if stringutils.IsEmptyOrNull(u.FirstName) {
		return ErrMissingFirstName
	}

	if stringutils.IsEmptyOrNull(u.LastName) {
		return ErrMissingLastName
	}

	if stringutils.IsEmptyOrNull(u.RoleId) {
		return ErrMissingRole
	}

	if requestType == SAVE_REQUEST_TYPE {
		if stringutils.IsEmptyOrNull(u.Email) {
			return ErrMissingEmail
		}

		if stringutils.IsEmptyOrNull(u.Password) {
			return ErrMissingPassword
		}

		if !stringutils.IsValidEmail(u.Email) {
			return ErrInvalidEmail
		}

		// TODO: Make a password policy to verify password
		if len(u.Password) < 5 {
			return ErrInvalidPassword
		}
	}

	return nil
}

func (u *UserDomain) ToUserEntity(requestType int) User {
	user := User{}
	user.Id = u.Id
	user.FirstName = u.FirstName
	user.LastName = u.LastName
	user.ProfilePictureUrl = u.ProfilePictureUrl
	user.RoleId = u.RoleId
	user.DeletedAt = ""
	user.IsActive = true
	user.IsDeleted = false

	if requestType == SAVE_REQUEST_TYPE {
		user.Email = u.Email
		user.PasswordHash = u.Password
		user.CreatedAt = dateutils.GetNow().String()
		user.UpdatedAt = user.CreatedAt
	}

	if requestType == UPDATE_REQUEST_TYPE {
		user.UpdatedAt = dateutils.GetNow().String()
	}

	return user
}

func (u *User) ToUserDomain() *UserDomain {
	domain := UserDomain{}
	domain.Id = u.Id
	domain.FirstName = u.FirstName
	domain.LastName = u.LastName
	domain.Email = u.Email
	domain.RoleId = u.RoleId
	domain.ProfilePictureUrl = u.ProfilePictureUrl
	domain.CreatedAt = u.CreatedAt
	domain.UpdatedAt = u.UpdatedAt
	domain.DeletedAt = u.DeletedAt
	domain.IsActive = u.IsActive
	domain.IsDeleted = u.IsDeleted
	return &domain
}

func (ulist UserList) ToUserDomainList() UserDomainList {
	var list UserDomainList
	for _, val := range ulist {
		list = append(list, *val.ToUserDomain())
	}
	return list
}
