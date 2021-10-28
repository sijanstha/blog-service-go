package comment

import "errors"

var (
	ErrInvalidReview = errors.New("invalid comment")
	ErrReviewMissing = errors.New("comment missing")
)

const (
	MIN_REVIEW_LENGTH = 5
	MAX_REVIEW_LENGTH = 300
)

type Comment struct {
	Id        string `json:"id"`
	Review    string `josn:"review"`
	IsActive  bool   `json:"active"`
	IsDeleted bool   `json:"deleted"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

func (c *Comment) Validate() error {
	if c.Review == "" || len(c.Review) == 0 {
		return ErrReviewMissing
	}

	if len(c.Review) < MIN_REVIEW_LENGTH || len(c.Review) > MAX_REVIEW_LENGTH {
		return ErrInvalidReview
	}

	return nil
}
