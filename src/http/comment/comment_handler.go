package commenthandler

import (
	"net/http"

	cdomain "github.com/blog-service/src/domain/comment"
	"github.com/blog-service/src/service/comment"
	dateutils "github.com/blog-service/src/utils/date"
	"github.com/gin-gonic/gin"
)

type ICommentHandler interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	GetById(c *gin.Context)
	Get(c *gin.Context)
	GetAll(c *gin.Context)
	GetAllWithPagination(c *gin.Context)
	Delete(c *gin.Context)
}

type commentHandler struct {
	commentService comment.ICommentService
}

func NewCommentHandler(commentService comment.ICommentService) ICommentHandler {
	return &commentHandler{commentService}
}

func (handler *commentHandler) Create(c *gin.Context) {
	var request cdomain.Comment
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid JSON body")
		return
	}

	result, err := handler.commentService.Save(&request)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (handler *commentHandler) Update(c *gin.Context) {
	var request cdomain.Comment
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid JSON body")
		return
	}

	result, err := handler.commentService.Update(&request)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusAccepted, result)
}

func (handler *commentHandler) GetById(c *gin.Context) {
	post, err := handler.commentService.FindById(c.Param("comment_id"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, post)
}

func (handler *commentHandler) Get(c *gin.Context) {
	var request cdomain.CommentFilter
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid JSON body")
		return
	}
	result, err := handler.commentService.Find(request)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusAccepted, result)
}

func (handler *commentHandler) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, handler.commentService.FindAll(c.Param("post_id")))
}

func (handler *commentHandler) GetAllWithPagination(c *gin.Context) {
	var request cdomain.CommentListFilter
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid JSON body")
		return
	}
	result, err := handler.commentService.FindAllWithPagination(request)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (handler *commentHandler) Delete(c *gin.Context) {
	err := handler.commentService.Delete(c.Param("comment_id"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	response := map[string]interface{}{
		"message":   "Deleted Successfully",
		"timestamp": dateutils.GetNow().Unix(),
	}
	c.JSON(http.StatusOK, response)
}
