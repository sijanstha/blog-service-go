package comment

import (
	"github.com/blog-service/src/domain/comment"
	commentrepo "github.com/blog-service/src/repository/comment"
	"github.com/blog-service/src/repository/post"
	dateutils "github.com/blog-service/src/utils/date"
	"github.com/blog-service/src/utils/errors"
	stringutils "github.com/blog-service/src/utils/string"
)

type ICommentService interface {
	Save(*comment.Comment) (*comment.Comment, *errors.RestErr)
	Update(*comment.Comment) (*comment.Comment, *errors.RestErr)
	FindById(string) (*comment.Comment, *errors.RestErr)
	Find(comment.CommentFilter) (*comment.Comment, *errors.RestErr)
	FindAll(string) []comment.Comment
	FindAllWithPagination(comment.CommentListFilter) (*comment.CommentPaginationDetails, *errors.RestErr)
	Delete(string) *errors.RestErr
}

type commentService struct {
	postRepository    post.IPostRepository
	commentRepository commentrepo.ICommentRepository
}

func NewCommentService(postRepository post.IPostRepository, commentRepository commentrepo.ICommentRepository) ICommentService {
	return &commentService{
		postRepository,
		commentRepository,
	}
}

func (c *commentService) Save(request *comment.Comment) (*comment.Comment, *errors.RestErr) {
	err := request.Validate()
	if err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}

	_, err = c.postRepository.FindById(request.PostId)
	if err != nil {
		return nil, errors.NewNotFoundError(err.Error())
	}

	request.Id = stringutils.GenerateUniqueId()
	request.CreatedAt = dateutils.GetNow().String()
	request.UpdatedAt = request.CreatedAt
	request.DeletedAt = ""
	request.IsActive = true
	request.IsDeleted = false

	result, err := c.commentRepository.Save(*request)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return result, nil
}

func (c *commentService) Update(request *comment.Comment) (*comment.Comment, *errors.RestErr) {
	if request.Id == "" || len(request.Id) <= 0 {
		return nil, errors.NewBadRequestError("comment id missing")
	}

	err := request.Validate()
	if err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}

	_, restErr := c.FindById(request.Id)
	if restErr != nil {
		return nil, restErr
	}

	request.UpdatedAt = dateutils.GetNow().String()
	request.IsActive = true
	request.IsDeleted = false

	result, err := c.commentRepository.Update(*request)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return result, nil
}

func (c *commentService) FindById(id string) (*comment.Comment, *errors.RestErr) {
	result, err := c.commentRepository.FindById(id)
	if err != nil {
		if err == commentrepo.ErrCommentNotFound {
			return nil, errors.NewNotFoundError(err.Error())
		}
		return nil, errors.NewInternalServerError(err.Error())
	}

	return result, nil
}

func (c *commentService) Find(filter comment.CommentFilter) (*comment.Comment, *errors.RestErr) {
	if err := filter.Validate(); err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}
	result, err := c.commentRepository.Find(filter)
	if err != nil {
		if err == commentrepo.ErrCommentNotFound {
			return nil, errors.NewNotFoundError(err.Error())
		}
		return nil, errors.NewInternalServerError(err.Error())
	}

	return result, nil
}

func (c *commentService) FindAll(postId string) []comment.Comment {
	return c.commentRepository.FindAll(comment.CommentFilter{PostId: postId})
}

func (c *commentService) FindAllWithPagination(filter comment.CommentListFilter) (*comment.CommentPaginationDetails, *errors.RestErr) {
	if err := filter.Validate(); err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}
	return c.commentRepository.FindAllWithPagination(filter), nil
}

func (c *commentService) Delete(id string) *errors.RestErr {
	req, restErr := c.FindById(id)
	if restErr != nil {
		return restErr
	}

	err := c.commentRepository.Delete(req)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
