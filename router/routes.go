package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"

	"restful.api.e-commerce.golang/infra/database"
	_memberHandlerHttp "restful.api.e-commerce.golang/modules/member/delivery/http"
	_mongoMemberRepo "restful.api.e-commerce.golang/modules/member/repository/mongo"
	_redisMemberRepo "restful.api.e-commerce.golang/modules/member/repository/redis"
	_memberUsecase "restful.api.e-commerce.golang/modules/member/usecase"

	_productHandlerHttp "restful.api.e-commerce.golang/modules/product/delivery/http"
	_mongoProductRepo "restful.api.e-commerce.golang/modules/product/repository/mongo"
	_productUsecase "restful.api.e-commerce.golang/modules/product/usecase"

	_orderHandlerHttp "restful.api.e-commerce.golang/modules/order/delivery/http"
	_mongoOrderRepo "restful.api.e-commerce.golang/modules/order/repository/mongo"
	_orderUsecase "restful.api.e-commerce.golang/modules/order/usecase"

	_redisAuthRepo "restful.api.e-commerce.golang/modules/auth/repository/redis"
	_authUsecase "restful.api.e-commerce.golang/modules/auth/usecase"
)

var (
	mongoDB *mongo.Database
	redisDB *redis.Client
)

func init() {
	mongoDB = database.ConnectMongoDB()
	redisDB = database.ConnectRedis()
}

func Index(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "server is alive",
		})
	})

	redisAuthRepo := _redisAuthRepo.NewRedisAuthRepo(redisDB)
	authUsecase := _authUsecase.NewAuthUsecase(redisAuthRepo)

	mongoMemberRepo := _mongoMemberRepo.NewMongoMemberRepo(mongoDB)
	redisMemberRepo := _redisMemberRepo.NewRedisMemberRepo(redisDB)
	memberUsecase := _memberUsecase.NewMemberUsecase(mongoMemberRepo, redisMemberRepo)
	_memberHandlerHttp.NewMemberHandler(r, memberUsecase)

	mongoProductRepo := _mongoProductRepo.NewMongoProductRepo(mongoDB)
	productUsecase := _productUsecase.NewProductUsecase(mongoProductRepo)
	_productHandlerHttp.NewProductHandler(r.Group("product"), productUsecase, authUsecase)

	mongoOrderRepo := _mongoOrderRepo.NewMongoOrderRepo(mongoDB)
	orderUsecase := _orderUsecase.NewOrderUsecase(mongoOrderRepo, mongoProductRepo)
	_orderHandlerHttp.NewOrderHandler(r.Group("order"), orderUsecase, authUsecase)

}
