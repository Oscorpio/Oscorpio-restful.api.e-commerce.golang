package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"restful.api.e-commerce.golang/domain"
	"restful.api.e-commerce.golang/utils"
)

type productHandler struct {
	productUsecase domain.ProductUsecase
}

func NewProductHandler(r *gin.RouterGroup, dp domain.ProductUsecase) {
	handler := &productHandler{
		productUsecase: dp,
	}

	r.GET("/:productId", handler.ListProductById)
	r.GET("/", handler.ListProducts)
	r.POST("/", handler.StoreProduct)
	r.POST("/uploadImage", handler.UploadImage)
	r.POST("/unit", handler.StoreUnitStock)
}

func (p *productHandler) UploadImage(ctx *gin.Context) {
	image, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	id, _ := ctx.GetPostForm("productId")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": domain.ErrParamInput.Error(),
		})
	}

	r, err := p.productUsecase.StoreImage(ctx, image, oid)
	if err != nil {
		ctx.JSON(utils.GetHttpStatus(err), err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"_id": r,
	})
}

func (p *productHandler) StoreProduct(ctx *gin.Context) {
	params := &domain.Product{}

	err := ctx.ShouldBind(params)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": domain.ErrParamInput.Error(),
		})
		return
	}

	err = p.productUsecase.CreateProduct(ctx, params)
	if err != nil {
		ctx.JSON(utils.GetHttpStatus(err), gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})
}

func (p *productHandler) ListProducts(ctx *gin.Context) {
	r, err := p.productUsecase.ListProducts(ctx)
	if err != nil {
		ctx.JSON(utils.GetHttpStatus(err), gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, r)
}

func (p *productHandler) StoreUnitStock(ctx *gin.Context) {
	params := &domain.UnitProduct{}
	err := ctx.ShouldBind(params)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": domain.ErrParamInput.Error(),
		})
		return
	}

	err = p.productUsecase.StoreUnitStock(ctx, params)
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

func (p *productHandler) ListProductById(ctx *gin.Context) {
	id := ctx.Param("productId")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": err.Error(),
		})
		return
	}

	r, err := p.productUsecase.ListProductById(ctx, oid)
	if err != nil {
		ctx.JSON(utils.GetHttpStatus(err), gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, r)
}
