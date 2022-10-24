package main

import (
	"fmt"
	"os"
	http_controllers "user/app/controllers"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	listeningPort := os.Getenv("SERVICE_PORT")

	router := http_controllers.InitControllers()

	if err := router.Run(":" + listeningPort); err != nil {
		fmt.Println(err)
	}
}
