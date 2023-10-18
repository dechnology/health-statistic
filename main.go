package main

import (
	"context"
	"log"
	"os"

	"github.com/eesoymilk/health-statistic-api/db"
	_ "github.com/eesoymilk/health-statistic-api/docs"
	"github.com/eesoymilk/health-statistic-api/router"
	"github.com/joho/godotenv"
)

// @Title		Web3 - 健康資料公鏈API開發文件
// @Version	1.0
// @Description.markdown
// @Host						health-statistic.dechnology.com.tw
// @Schemes					https
// @BasePath					/api/v1
// @SecurityDefinitions.apikey	BearerAuth
// @In							header
// @Name						Authorization
func main() {
	// Loading environment variables when not using a docker container
	if os.Getenv("RUNNING_IN_DOCKER") == "" {
		log.Print("Not running in a docker container. Loading .env file...")
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	entClient := db.New()
	defer entClient.Close()

	// migration code: perform it only onces is enough
	if err := db.Migrate(context.Background(), entClient); err != nil {
		log.Print(err.Error())
	}

	r := router.New(entClient)
	r.Run()
}
