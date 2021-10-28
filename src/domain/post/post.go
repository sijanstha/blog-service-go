package post

import (
	"errors"
)

var (
	ErrInvalidTile        = errors.New("invalid post title")
	ErrInvalidDescription = errors.New("invalid post description")
	ErrTitleMissing       = errors.New("title missing")
	ErrDescriptionMissing = errors.New("description missing")
)

const (
	MIN_TITLE_LENGTH       = 5
	MAX_TITLE_LENGTH       = 100
	MIN_DESCRIPTION_LENGTH = 50
	MAX_DESCRIPTION_LENGTH = 1000
)

type Post struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsActive    bool   `json:"active"`
	IsDeleted   bool   `json:"deleted"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}

type PostPaginationDetails struct {
	Size  int    `json:"pageSize"`
	Page  int    `json:"page"`
	Total int    `json:"totalRecords"`
	Data  []Post `json:"data"`
}

func (p *Post) Validate() error {
	if p.Title == "" || len(p.Title) == 0 {
		return ErrTitleMissing
	}

	if p.Description == "" || len(p.Description) == 0 {
		return ErrDescriptionMissing
	}

	if len(p.Title) < MIN_TITLE_LENGTH || len(p.Title) > MAX_TITLE_LENGTH {
		return ErrInvalidTile
	}

	if len(p.Description) < MIN_DESCRIPTION_LENGTH || len(p.Description) > MAX_DESCRIPTION_LENGTH {
		return ErrInvalidDescription
	}

	return nil
}
