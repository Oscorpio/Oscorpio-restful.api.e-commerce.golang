package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"restful.api.e-commerce.golang/domain"
	"restful.api.e-commerce.golang/utils"
)

type productHandler struct {
	productUsecase domain.ProductUsecase
	authUsecase    domain.AuthUsecase
}

func NewProductHandler(r *gin.RouterGroup, dp domain.ProductUsecase, da domain.AuthUsecase) {
	handler := &productHandler{
		productUsecase: dp,
		authUsecase:    da,
	}

	r.GET("/detail/:productId", handler.ListProductById)
	r.GET("/", handler.ListProducts)
	r.POST("/", handler.StoreProduct)
	r.POST("/uploadImage", handler.UploadImage)
	r.POST("/unit", handler.StoreDetail)
}

func (p *productHandler) UploadImage(ctx *gin.Context) {
	err := p.authUsecase.ValidateToken(ctx, ctx.GetHeader("Authentication"))
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": domain.ErrForbidden.Error(),
		})
		return
	}
	image, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	fileNameSplit := strings.Split(image.Filename, ".")
	fileType := fileNameSplit[len(fileNameSplit)-1]
	if fileType != "png" && fileType != "jpg" && fileType != "jpeg" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": domain.ErrParamInput.Error(),
		})
		return
	}

	id, _ := ctx.GetPostForm("productId")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": domain.ErrParamInput.Error(),
		})
	}

	err = p.productUsecase.StoreImage(ctx, image, oid)
	if err != nil {
		ctx.JSON(utils.GetHttpStatus(err), err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"msg": "success",
	})
}

func (p *productHandler) StoreProduct(ctx *gin.Context) {
	err := p.authUsecase.ValidateToken(ctx, ctx.GetHeader("Authentication"))
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": domain.ErrForbidden.Error(),
		})
		return
	}

	params := &domain.Product{}

	err = ctx.ShouldBind(params)
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

func (p *productHandler) StoreDetail(ctx *gin.Context) {
	err := p.authUsecase.ValidateToken(ctx, ctx.GetHeader("Authentication"))
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": domain.ErrForbidden.Error(),
		})
		return
	}

	params := &domain.Detail{}
	err = ctx.ShouldBind(params)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": domain.ErrParamInput.Error(),
		})
		return
	}

	err = p.productUsecase.StoreDetail(ctx, params)
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
