package post

import (
	"reflect"

	"github.com/blog-service/src/utils/errors"
	stringutils "github.com/blog-service/src/utils/string"
)

type PostFilter struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Active  *bool  `json:"active"`
	Deleted *bool  `json:"deleted"`
}

type PostListFilter struct {
	Filter    PostFilter `json:"filter"`
	CreatedAt string     `json:"created_at"`
	Limit     int64      `json:"limit" default:"10"`
	Page      int64      `json:"page" default:"1"`
	Sort      string     `json:"sort" default:"asc"`
	SortBy    string     `json:"sortBy" default:"title"`
}

func (filter *PostFilter) Validate() error {
	if filter.Id == "" && filter.Title == "" {
		return errors.ErrFilterConditionMissing
	}

	return nil
}

func (filter *PostListFilter) Validate() error {
	typ := reflect.TypeOf(*filter)
	if filter.Limit == 0 {
		field, _ := typ.FieldByName("Limit")
		filter.Limit = stringutils.ParseInteger(field.Tag.Get("default"))
	}
	if filter.Page == 0 {
		field, _ := typ.FieldByName("Page")
		filter.Page = stringutils.ParseInteger(field.Tag.Get("default"))
	}
	if filter.Sort == "" {
		field, _ := typ.FieldByName("Sort")
		filter.Sort = field.Tag.Get("default")
	}
	if filter.SortBy == "" {
		field, _ := typ.FieldByName("SortBy")
		filter.SortBy = field.Tag.Get("default")
	}
	if filter.SortBy != "" && len(filter.SortBy) > 0 {
		if !errors.ValidSortingField[filter.SortBy] {
			return errors.ErrInvalidSortingField
		}
	}
	if filter.Sort != "" && len(filter.Sort) > 0 {
		if !errors.ValidSortingOrder[filter.Sort] {
			return errors.ErrInvalidSortingOrder
		}
	}
	return nil
}
