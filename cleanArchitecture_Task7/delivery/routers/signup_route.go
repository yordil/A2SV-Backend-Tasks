package routers

import (
	"fmt"
	"task7/Delivery/controllers"
	"task7/bootstrap"
	"task7/repository"
	"task7/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewSignUPRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Client, router *gin.RouterGroup) {
	

	userRepo := repository.NewUserRepositoryImpl(db ,  env.DBName, env.UserCollection)
	userUseCase := usecase.NewUserUsecase(userRepo)
	userController := controllers.NewUserController(userUseCase)
	fmt.Println(userController)
	router.POST("/signup", userController.SignUp)


}