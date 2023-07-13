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
	"log"
	"os"

	"github.com/eesoymilk/health-statistic-api/router"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("RUNNING_IN_DOCKER") == "" {
		log.Print("Not running in a docker container. Loading .env file...")
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	r := router.New()
	r.Run()
}
