package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"restful.api.e-commerce.golang/infra/database"
	_memberHandlerHttp "restful.api.e-commerce.golang/modules/member/delivery/http"
	_memberRepo "restful.api.e-commerce.golang/modules/member/repository/mongo"
	_memberUsecase "restful.api.e-commerce.golang/modules/member/usecase"
)

var db *mongo.Database

func init() {
	db = database.ConnectMongoDB()
}

func Index(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "server is alive",
		})
	})

	memberRepo := _memberRepo.NewMongoMemberRepo(db)
	memberUsecase := _memberUsecase.NewMemberUsecase(memberRepo)
	_memberHandlerHttp.NewMemberHandler(r, memberUsecase)

}
