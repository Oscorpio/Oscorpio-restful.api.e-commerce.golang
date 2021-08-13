package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"restful.api.e-commerce.golang/domain"
	"restful.api.e-commerce.golang/utils"
)

type memberHandler struct {
	memberUsecase domain.MemberUsecase
}

func NewMemberHandler(r *gin.RouterGroup, dm domain.MemberUsecase) {
	handler := &memberHandler{
		memberUsecase: dm,
	}

	r.POST("/signUp", handler.CreateUser)
	r.POST("/login", handler.Login)
}

func (m *memberHandler) CreateUser(c *gin.Context) {
	params := domain.User{}

	bindErr := c.ShouldBind(&params)
	if bindErr != nil {
		c.JSON(http.StatusUnprocessableEntity, bindErr.Error())
		return
	}

	err := m.memberUsecase.CreateUser(c, &params)
	if err != nil {
		c.JSON(utils.GetHttpStatus(err), err.Error())
		return
	}

	c.JSON(http.StatusCreated, "success")
}

func (m *memberHandler) Login(c *gin.Context) {
	type v struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	params := v{}

	bindErr := c.ShouldBind(&params)
	if bindErr != nil {
		c.JSON(http.StatusUnprocessableEntity, bindErr.Error())
		return
	}

	token, err := m.memberUsecase.Login(c, params.Email, params.Password)
	if err != nil {
		c.JSON(utils.GetHttpStatus(err), err.Error())
		return
	}

	c.JSON(http.StatusOK, token)
}
