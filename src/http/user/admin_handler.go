package userhandler

import (
	"net/http"

	userdomain "github.com/blog-service/src/domain/user"
	"github.com/blog-service/src/service/user"
	dateutils "github.com/blog-service/src/utils/date"
	"github.com/gin-gonic/gin"
)

type IAdminUserAdminHandler interface {
	GetAllUsers(c *gin.Context)
	GetAllUsersWithPagination(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userAdminHandler struct {
	userService user.IUserService
}

func NewAdminuserAdminHandler(userService user.IUserService) IAdminUserAdminHandler {
	return &userAdminHandler{userService}
}

func (handler *userAdminHandler) GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, handler.userService.FindAll())
}

func (handler *userAdminHandler) GetAllUsersWithPagination(c *gin.Context) {
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

func (handler *userAdminHandler) DeleteUser(c *gin.Context) {
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
