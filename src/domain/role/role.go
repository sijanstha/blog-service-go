package role

import (
	"errors"

	stringutils "github.com/blog-service/src/utils/string"
)

var (
	ErrMissingRoleName = errors.New("role name missing")
	ErrInvalidRoleName = errors.New("invalid rome name")
)

type Role struct {
	Id          string `json:"id"`
	RoleName    string `json:"roleName"`
	DisplayName string `json:"displayName"`
	IsActive    bool   `json:"active"`
	IsDeleted   bool   `json:"deleted"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	DeletedAt   string `json:"deletedAt"`
}

func (role *Role) Validate() error {
	if stringutils.IsEmptyOrNull(role.RoleName) {
		return ErrMissingRoleName
	}

	if len(role.RoleName) > 50 {
		return ErrInvalidRoleName
	}

	return nil
}
