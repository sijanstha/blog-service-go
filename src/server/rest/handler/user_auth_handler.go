package handler

import (
	"net/http"

	"github.com/blog-service/src/service"
	"github.com/blog-service/src/utils/errors"
	"github.com/gin-gonic/gin"

	userdomain "github.com/blog-service/src/domain/user"
)

type IUserAuthHandler interface {
	Register(*gin.Context)
	Login(*gin.Context)
}

type userAuthHandler struct {
	userAuthService service.IUserAuthService
}

func NewUserAuthHandler(userAuthService service.IUserAuthService) IUserAuthHandler {
	return &userAuthHandler{userAuthService}
}

func (handler *userAuthHandler) Register(c *gin.Context) {
	var request userdomain.UserDomain

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("Invalid JSON body"))
		return
	}

	result, err := handler.userAuthService.Register(&request)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, result)

}

func (handler *userAuthHandler) Login(c *gin.Context) {
	var request userdomain.UserLoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("Invalid JSON body"))
		return
	}

	result, err := handler.userAuthService.Login(&request)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}
