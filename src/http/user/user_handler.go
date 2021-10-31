package userhandler

import (
	"net/http"

	userdomain "github.com/blog-service/src/domain/user"
	"github.com/blog-service/src/service/user"
	dateutils "github.com/blog-service/src/utils/date"
	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	GetById(c *gin.Context)
	Get(c *gin.Context)
	GetAll(c *gin.Context)
	GetAllWithPagination(c *gin.Context)
	Delete(c *gin.Context)
}

type userHandler struct {
	userService user.IUserService
}

func NewUserHandler(userService user.IUserService) IUserHandler {
	return &userHandler{userService}
}

func (handler *userHandler) Create(c *gin.Context) {
	var request userdomain.UserDomain
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid JSON body")
		return
	}

	result, err := handler.userService.Save(&request)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (handler *userHandler) Update(c *gin.Context) {
	var request userdomain.UserDomain
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid JSON body")
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
		c.JSON(http.StatusBadRequest, "Invalid JSON body")
		return
	}
	result, err := handler.userService.Find(request)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusAccepted, result)
}

func (handler *userHandler) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, handler.userService.FindAll())
}

func (handler *userHandler) GetAllWithPagination(c *gin.Context) {
	var request userdomain.UserListFilter
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid JSON body")
		return
	}
	result, err := handler.userService.FindAllWithPagination(request)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (handler *userHandler) Delete(c *gin.Context) {
	err := handler.userService.Delete(c.Param("user_id"))
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
