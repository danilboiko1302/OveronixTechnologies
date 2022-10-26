package main

import (
	"fmt"
	"os"
	http_controllers "user/app/controllers"
	"user/app/db/queries"
	_ "user/docs"

	"github.com/joho/godotenv"
)

// @title           OveronixTechnologies
// @version         1.0
// @description     User CRUD.
// @termsOfService  http://swagger.io/terms/

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /users

// @securityDefinitions.basic  BasicAuth
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
