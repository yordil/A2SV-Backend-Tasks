package main

import (
	"fmt"
	"task7/bootstrap"
	"task7/delivery/routers"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
    
    fmt.Println("here worked")
    
    app := bootstrap.App()
    env := app.Env

    db := app.Mongo

    timeout := time.Duration(env.ContextTimeout) * time.Second

    gin := gin.Default()

    routers.SetupRouter(env , timeout , &db , gin)
    
    gin.Run(env.ServerAddress)
}
