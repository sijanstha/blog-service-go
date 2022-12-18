package handler

import (
	"net/http"

	userdomain "github.com/blog-service/src/domain/user"
	"github.com/blog-service/src/service"
	"github.com/blog-service/src/utils/errors"
	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	Update(c *gin.Context)
	GetById(c *gin.Context)
	Get(c *gin.Context)
}

type userHandler struct {
	userService service.IUserService
}

func NewUserHandler(userService service.IUserService) IUserHandler {
	return &userHandler{userService}
}

func (handler *userHandler) Update(c *gin.Context) {
	var request userdomain.UserDomain
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("Invalid JSON body"))
		return
	}

	result, err := handler.userService.Update(&request)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusAccepted, result)
}

func (handler *userHandler) GetById(c *gin.Context) {
	post, err := handler.userService.FindById(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, post)
}

func (handler *userHandler) Get(c *gin.Context) {
	var request userdomain.UserFilter
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("Invalid JSON body"))
		return
	}
	result, err := handler.userService.Find(request)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusAccepted, result)
}
