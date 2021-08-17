package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"restful.api.e-commerce.golang/domain"
	"restful.api.e-commerce.golang/utils"
)

type orderHandler struct {
	orderUsecase domain.OrderUsecase
	authUsecase  domain.AuthUsecase
}

func NewOrderHandler(r *gin.RouterGroup, do domain.OrderUsecase, da domain.AuthUsecase) {
	handler := &orderHandler{
		orderUsecase: do,
		authUsecase:  da,
	}

	r.POST("/", handler.CreateOrder)
	r.GET("/:owner", handler.ListOrder)
}

func (o *orderHandler) CreateOrder(ctx *gin.Context) {
	err := o.authUsecase.ValidateToken(ctx, ctx.GetHeader("Authentication"))
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": domain.ErrForbidden.Error(),
		})
		return
	}

	order := &domain.Order{}
	err = ctx.ShouldBind(order)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": domain.ErrParamInput.Error(),
		})
		return
	}
	err = o.orderUsecase.CreateOrder(ctx, order)
	if err != nil {
		ctx.JSON(utils.GetHttpStatus(err), gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"msg": "success",
	})
}

func (o *orderHandler) ListOrder(ctx *gin.Context) {
	err := o.authUsecase.ValidateToken(ctx, ctx.GetHeader("Authentication"))
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": domain.ErrForbidden.Error(),
		})
		return
	}

	param := ctx.Param("owner")
	if param == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": domain.ErrParamInput.Error(),
		})
	}
	r, err := o.orderUsecase.ListOrder(ctx, param)
	if err != nil {
		ctx.JSON(utils.GetHttpStatus(err), gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, r)
}
