package http

import (
	"net/http"
	"strings"

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
	r.POST("/logout", handler.Logout)
}

func (m *memberHandler) CreateUser(c *gin.Context) {
	params := domain.User{}

	bindErr := c.ShouldBind(&params)
	if bindErr != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": domain.ErrParamInput.Error(),
		})
		return
	}

	err := m.memberUsecase.CreateUser(c, &params)
	if err != nil {
		c.JSON(utils.GetHttpStatus(err), gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"msg": "success",
	})
}

func (m *memberHandler) Login(c *gin.Context) {
	type v struct {
		Email    string `json:"email" bson:"email" binding:"required,email"`
		Password string `json:"password" bson:"password" binding:"required"`
	}
	params := &v{}

	bindErr := c.ShouldBind(&params)
	if bindErr != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": domain.ErrParamInput.Error(),
		})
		return
	}

	token, err := m.memberUsecase.Login(c, params.Email, params.Password)
	if err != nil {
		c.JSON(utils.GetHttpStatus(err), gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, token)
}

func (m *memberHandler) Logout(ctx *gin.Context) {
	header := ctx.GetHeader("Authentication")
	if len(header) <= 0 {
		ctx.JSON(http.StatusForbidden, domain.ErrForbidden.Error())
	}

	token := strings.Split(header, "Bearer ")[1]
	if len(token) <= 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": domain.ErrParamInput.Error(),
		})
		return
	}

	err := m.memberUsecase.Logout(ctx, token)
	if err != nil {
		ctx.JSON(utils.GetHttpStatus(err), gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{
		"msg": "success",
	})
}
