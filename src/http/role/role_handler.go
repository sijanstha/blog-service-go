package rolehandler

import (
	"net/http"

	rdomain "github.com/blog-service/src/domain/role"
	"github.com/blog-service/src/service/role"
	dateutils "github.com/blog-service/src/utils/date"
	"github.com/gin-gonic/gin"
)

type IRoleHandler interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	GetById(c *gin.Context)
	Get(c *gin.Context)
	GetAll(c *gin.Context)
	Delete(c *gin.Context)
}

type roleHandler struct {
	roleService role.IRoleService
}

func NewRoleHandler(roleService role.IRoleService) IRoleHandler {
	return &roleHandler{roleService}
}

func (handler *roleHandler) Create(c *gin.Context) {
	var request rdomain.Role
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid JSON body")
		return
	}

	result, err := handler.roleService.Save(&request)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (handler *roleHandler) Update(c *gin.Context) {
	var request rdomain.Role
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid JSON body")
		return
	}

	result, err := handler.roleService.Update(&request)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusAccepted, result)
}

func (handler *roleHandler) GetById(c *gin.Context) {
	post, err := handler.roleService.FindById(c.Param("role_id"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, post)
}

func (handler *roleHandler) Get(c *gin.Context) {
	var request rdomain.RoleFilter
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid JSON body")
		return
	}
	result, err := handler.roleService.Find(request)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusAccepted, result)
}

func (handler *roleHandler) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, handler.roleService.FindAll())
}

func (handler *roleHandler) Delete(c *gin.Context) {
	err := handler.roleService.Delete(c.Param("role_id"))
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
