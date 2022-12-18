package handler

import (
	"github.com/blog-service/src/server/grpc/pb"
	"github.com/blog-service/src/service"
)

type GrpcBeanManager struct {
	*service.BeanFactory
}

func NewGrpcBeanManager(s *service.BeanFactory) *GrpcBeanManager {
	return &GrpcBeanManager{s}
}

func (rm GrpcBeanManager) NewPostServiceServerHandler() pb.PostServiceServer {
	return &postServiceServerHandler{postService: rm.GetPostService()}
}
