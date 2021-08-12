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
