package posthandler

import (
	"net/http"

	pdomain "github.com/blog-service/src/domain/post"
	"github.com/blog-service/src/service/post"
	"github.com/gin-gonic/gin"
)

type PostHandler interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	GetById(c *gin.Context)
	Get(c *gin.Context)
	GetAll(c *gin.Context)
	GetAllWithPagination(c *gin.Context)
}

type postHandler struct {
	postService post.PostService
}

func NewPostHandler(postService post.PostService) PostHandler {
	return &postHandler{postService}
}

func (handler *postHandler) Create(c *gin.Context) {
	var request pdomain.Post
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid JSON body")
		return
	}

	result, err := handler.postService.Save(&request)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (handler *postHandler) Update(c *gin.Context) {
	var request pdomain.Post
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid JSON body")
		return
	}

	result, err := handler.postService.Update(&request)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusAccepted, result)
}

func (handler *postHandler) GetById(c *gin.Context) {
	post, err := handler.postService.FindById(c.Param("post_id"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, post)
}

func (handler *postHandler) Get(c *gin.Context) {
	var request pdomain.PostFilter
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid JSON body")
		return
	}
	result, err := handler.postService.Find(request)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusAccepted, result)
}

func (handler *postHandler) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, handler.postService.FindAll())
}

func (handler *postHandler) GetAllWithPagination(c *gin.Context) {
	var request pdomain.PostListFilter
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid JSON body")
		return
	}
	result, err := handler.postService.FindAllWithPagination(request)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
