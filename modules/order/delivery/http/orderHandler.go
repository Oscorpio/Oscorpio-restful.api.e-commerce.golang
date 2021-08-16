package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"restful.api.e-commerce.golang/domain"
	"restful.api.e-commerce.golang/utils"
)

type orderHandler struct {
	orderUsecase domain.OrderUsecase
}

func NewOrderHandler(r *gin.RouterGroup, do domain.OrderUsecase) {
	handler := &orderHandler{
		orderUsecase: do,
	}

	r.POST("/", handler.CreateOrder)
	r.GET("/:owner", handler.ListOrder)
}

func (o *orderHandler) CreateOrder(ctx *gin.Context) {
	order := &domain.Order{}
	err := ctx.ShouldBind(order)
	fmt.Println("====>")
	if err != nil {
		fmt.Println("====>", err)
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": domain.ErrParamInput.Error(),
		})
		return
	}
	fmt.Println("====>")
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
