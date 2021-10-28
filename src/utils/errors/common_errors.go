package errors

import "errors"

var (
	ErrFilterConditionMissing = errors.New("provide at least one filter condition")
	ErrInvalidSortingField    = errors.New("invalid sorting field")
	ErrInvalidSortingOrder    = errors.New("invalid sorting order")
	ErrRowInsertFailed        = errors.New("cannot perform insert query at this moment")
	ErrInvalidQuery           = errors.New("invalid query")

	ValidSortingField = map[string]bool{
		"title":       true,
		"created_at":  true,
		"updated_at":  true,
		"description": true,
		"id":          true,
	}
	ValidSortingOrder = map[string]bool{
		"asc":  true,
		"desc": true,
	}
)
