package role

import "github.com/blog-service/src/utils/errors"

type RoleFilter struct {
	Id       string `json:"id"`
	RoleName string `json:"roleName"`
	Active   *bool  `json:"active"`
	Deleted  *bool  `json:"deleted"`
}

func (filter *RoleFilter) Validate() error {
	if filter.Id == "" && filter.RoleName == "" {
		return errors.ErrFilterConditionMissing
	}

	return nil
}
