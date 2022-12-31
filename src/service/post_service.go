package service

import (
	"context"
	"encoding/json"

	"github.com/blog-service/src/config"
	"github.com/blog-service/src/domain"
	"github.com/blog-service/src/domain/post"
	"github.com/blog-service/src/repository"
	dateutils "github.com/blog-service/src/utils/date"
	"github.com/blog-service/src/utils/errors"
	stringutils "github.com/blog-service/src/utils/string"
)

type IPostService interface {
	Save(*post.Post) (*post.Post, *errors.RestErr)
	Update(*post.Post) (*post.Post, *errors.RestErr)
	FindById(string) (*post.Post, *errors.RestErr)
	Find(post.PostFilter) (*post.Post, *errors.RestErr)
	FindAll() []post.Post
	FindAllWithPagination(post.PostListFilter) (*post.PostPaginationDetails, *errors.RestErr)
	Delete(string) *errors.RestErr
}

type postService struct {
	postRepo        repository.IPostRepository
	producer        *config.KafkaProducer
	postChangeTopic string
}

func NewPostService(postRepo repository.IPostRepository, kafkaProducer *config.KafkaProducer, postChangeTopic string) IPostService {
	return &postService{
		postRepo:        postRepo,
		producer:        kafkaProducer,
		postChangeTopic: postChangeTopic,
	}
}

func NewTestPostService(postRepo repository.IPostRepository) IPostService {
	return &postService{
		postRepo: postRepo,
	}
}

func (s *postService) Save(request *post.Post) (*post.Post, *errors.RestErr) {
	err := request.Validate()
	if err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}

	request.Id = stringutils.GenerateUniqueId()
	request.CreatedAt = dateutils.GetNow().String()
	request.UpdatedAt = request.CreatedAt
	request.DeletedAt = ""
	request.IsActive = true
	request.IsDeleted = false

	result, err := s.postRepo.Save(request)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	go func() {
		if s.producer != nil {
			msg := domain.Message{
				ChangeType: domain.CREATE,
				Body:       result,
			}
			stringJson, _ := json.Marshal(msg)
			s.producer.Produce(context.Background(), s.postChangeTopic, stringJson)
		}
	}()
	return result, nil
}

func (s *postService) Update(request *post.Post) (*post.Post, *errors.RestErr) {
	if request.Id == "" || len(request.Id) <= 0 {
		return nil, errors.NewBadRequestError("post id missing")
	}

	err := request.Validate()
	if err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}

	_, restErr := s.FindById(request.Id)
	if restErr != nil {
		return nil, restErr
	}

	request.UpdatedAt = dateutils.GetNow().String()
	request.IsActive = true
	request.IsDeleted = false

	result, err := s.postRepo.Update(request)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	go func() {
		if s.producer != nil {
			msg := domain.Message{
				ChangeType: domain.UPDATE,
				Body:       result,
			}
			stringJson, _ := json.Marshal(msg)
			s.producer.Produce(context.Background(), s.postChangeTopic, stringJson)
		}
	}()

	return result, nil
}

func (s *postService) FindById(id string) (*post.Post, *errors.RestErr) {
	result, err := s.postRepo.FindById(id)
	if err != nil {
		if err == repository.ErrPostNotFound {
			return nil, errors.NewNotFoundError(err.Error())
		}
		return nil, errors.NewInternalServerError(err.Error())
	}

	return result, nil
}

func (s *postService) Find(filter post.PostFilter) (*post.Post, *errors.RestErr) {
	if err := filter.Validate(); err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}
	result, err := s.postRepo.Find(filter)
	if err != nil {
		if err == repository.ErrPostNotFound {
			return nil, errors.NewNotFoundError(err.Error())
		}
		return nil, errors.NewInternalServerError(err.Error())
	}

	return result, nil
}

func (s *postService) FindAll() []post.Post {
	return s.postRepo.FindAll()
}

func (s *postService) FindAllWithPagination(filter post.PostListFilter) (*post.PostPaginationDetails, *errors.RestErr) {
	if err := filter.Validate(); err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}
	return s.postRepo.FindAllWithPagination(filter), nil
}

func (s *postService) Delete(id string) *errors.RestErr {
	req, restErr := s.FindById(id)
	if restErr != nil {
		return restErr
	}

	err := s.postRepo.Delete(req)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
