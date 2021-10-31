package user

import (
	"errors"

	stringutils "github.com/blog-service/src/utils/string"
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
}

func (u *UserDomain) Validate() error {
	if stringutils.IsEmptyOrNull(u.FirstName) {
		return ErrMissingFirstName
	}

	if stringutils.IsEmptyOrNull(u.LastName) {
		return ErrMissingLastName
	}

	if stringutils.IsEmptyOrNull(u.Email) {
		return ErrMissingEmail
	}

	if stringutils.IsEmptyOrNull(u.Password) {
		return ErrMissingPassword
	}

	if stringutils.IsEmptyOrNull(u.RoleId) {
		return ErrMissingRole
	}

	if !stringutils.IsValidEmail(u.Email) {
		return ErrInvalidEmail
	}

	// TODO: Make a password policy to verify password
	if len(u.Password) < 10 {
		return ErrInvalidPassword
	}

	return nil
}

func (u *UserDomain) ToEntity() User {
	user := User{}
	user.Id = u.Id
	user.FirstName = u.FirstName
	user.LastName = u.LastName
	user.Email = u.Email
	user.PasswordHash = u.Password
	user.ProfilePictureUrl = u.ProfilePictureUrl
	user.RoleId = u.RoleId
	return user
}
