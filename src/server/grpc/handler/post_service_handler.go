package handler

import (
	"context"

	"github.com/blog-service/src/domain/post"
	"github.com/blog-service/src/server/grpc/pb"
	"github.com/blog-service/src/service"
	"google.golang.org/grpc/codes"
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

	// TODO: Make generic error response struct
	if err != nil {
		return nil, status.Error(codes.Code(err.Code), err.Message)
	}

	// TODO: Make mapper util to map these from one another
	return &pb.PostResponse{
		Id:          res.Id,
		Title:       res.Title,
		Description: res.Description,
		IsActive:    res.IsActive,
		IsDeleted:   res.IsDeleted,
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
		DeletedAt:   res.DeletedAt,
	}, nil
}

func (psa postServiceServerHandler) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.PostResponse, error) {
	panic("to implement")
}

func (psa postServiceServerHandler) FindPost(ctx context.Context, req *pb.PostFilter) (*pb.PostResponse, error) {
	panic("to implement")
}
