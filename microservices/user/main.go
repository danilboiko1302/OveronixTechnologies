package main

import (
	"fmt"
	"os"
	http_controllers "user/app/controllers"
	"user/app/db/queries"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	if err := queries.Init(); err != nil {
		panic(err)
	}

	defer queries.SQLSession.Close()

	listeningPort := os.Getenv("SERVICE_PORT")

	router := http_controllers.InitControllers()

	if err := router.Run(":" + listeningPort); err != nil {
		fmt.Println(err)
	}
}
