package handler

import (
	"context"

	"github.com/blog-service/src/domain/post"
	"github.com/blog-service/src/server/grpc/mapper"
	"github.com/blog-service/src/server/grpc/pb"
	"github.com/blog-service/src/service"
	"google.golang.org/grpc/status"
)

type postServiceServerHandler struct {
	postService service.IPostService
	pb.UnimplementedPostServiceServer
}

func NewPostServiceServerHandler(postService service.IPostService) pb.PostServiceServer {
	return &postServiceServerHandler{postService: postService}
}

func (psa postServiceServerHandler) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.PostResponse, error) {
	res, err := psa.postService.Save(&post.Post{
		Title:       req.Title,
		Description: req.Description,
	})

	if err != nil {
		return nil, status.Error(mapper.ToGrpcErrorCode(err.Code), err.Message)
	}

	return res.ToPostResponse(), nil
}

func (psa postServiceServerHandler) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.PostResponse, error) {
	res, err := psa.postService.Update(&post.Post{
		Id:          req.Id,
		Title:       req.Title,
		Description: req.Description,
		IsActive:    req.IsActive,
		IsDeleted:   req.IsDeleted,
	})

	if err != nil {
		return nil, status.Error(mapper.ToGrpcErrorCode(err.Code), err.Message)
	}

	return res.ToPostResponse(), nil
}

func (psa postServiceServerHandler) FindPost(ctx context.Context, req *pb.PostFilter) (*pb.PostResponse, error) {
	res, err := psa.postService.Find(post.PostFilter{
		Id:      req.Id,
		Active:  &req.IsActive,
		Deleted: &req.IsDeleted,
	})

	if err != nil {
		return nil, status.Error(mapper.ToGrpcErrorCode(err.Code), err.Message)
	}

	return res.ToPostResponse(), nil
}
