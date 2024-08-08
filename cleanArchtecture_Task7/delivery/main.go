package main

import (

	"task7/delivery/routers"
)

func main() {
    

    router := routers.SetupRouter()
    router.Run(":8080")
}
