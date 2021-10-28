package comment

import (
	"errors"

	stringutils "github.com/blog-service/src/utils/string"
)

var (
	ErrInvalidReview = errors.New("invalid comment")
	ErrReviewMissing = errors.New("comment missing")
	ErrPostIdMissing = errors.New("post id missing")
)

const (
	MIN_REVIEW_LENGTH = 5
	MAX_REVIEW_LENGTH = 300
)

type Comment struct {
	Id        string `json:"id"`
	Review    string `json:"review"`
	PostId    string `json:"postId"`
	IsActive  bool   `json:"active"`
	IsDeleted bool   `json:"deleted"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	DeletedAt string `json:"deletedAt"`
}

type CommentPaginationDetails struct {
	Size  int       `json:"pageSize"`
	Page  int       `json:"page"`
	Total int       `json:"totalRecords"`
	Data  []Comment `json:"data"`
}

func (c *Comment) Validate() error {
	if stringutils.IsEmptyOrNull(c.Review) {
		return ErrReviewMissing
	}

	if len(c.Review) < MIN_REVIEW_LENGTH || len(c.Review) > MAX_REVIEW_LENGTH {
		return ErrInvalidReview
	}

	if stringutils.IsEmptyOrNull(c.PostId) {
		return ErrPostIdMissing
	}

	return nil
}
