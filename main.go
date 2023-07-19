package main

/**************  TODO  **************/
/*	1. Auth0 Integration			*/
/*	2. CRUD - User Information		*/
/*	3. CRUD - Questionnaire			*/
/*	4. CRUD - User Healthkit Data	*/
/*	5. CRUD - DeeGoo Links			*/
/*	6. CRUD - Reward Notification? 	*/
/*	7. Noitfication - Reward Keys?	*/
/*	8. Noitfication - Lottory Keys? */
/************************************/

import (
	"context"
	"log"
	"os"

	"github.com/eesoymilk/health-statistic-api/db"
	_ "github.com/eesoymilk/health-statistic-api/docs"
	"github.com/eesoymilk/health-statistic-api/router"
	"github.com/joho/godotenv"
)

// @title           Health Statistic API
// @version         1.0
// @description     This is a sample server for Health Statistic API.
// @license.name 	Apache 2.0
// @license.url 	http://www.apache.org/licenses/LICENSE-2.0.html
// @host      		localhost:8080
// @schemes			http https
// @BasePath  		/api/v1
func main() {
	// Loading environment variables when not using a docker container
	if os.Getenv("RUNNING_IN_DOCKER") == "" {
		log.Print("Not running in a docker container. Loading .env file...")
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	db := db.New()
	defer db.Close()

	// for migration
	if err := db.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	r := router.New(db)
	r.Run()
}
