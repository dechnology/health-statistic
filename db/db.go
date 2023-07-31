package db

import (
	"fmt"
	"log"
	"os"

	"github.com/eesoymilk/health-statistic-api/ent"
	_ "github.com/lib/pq"
)

func New() *ent.Client {
	var (
		postgresHost     = os.Getenv("POSTGRES_HOST")
		postgresPort     = os.Getenv("POSTGRES_PORT")
		postgresUser     = os.Getenv("POSTGRES_USER")
		postgresDb       = os.Getenv("POSTGRES_DB")
		postgresPassword = os.Getenv("POSTGRES_PASSWORD")
	)

	connectionStr := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		postgresHost,
		postgresPort,
		postgresUser,
		postgresDb,
		postgresPassword,
	)

	log.Println(connectionStr)
	//
	// drv, err := sql.Open("postgres", connectionStr)
	// if err != nil {
	// 	log.Fatalf("failed opening connection to postgres: %v", err)
	// }
	//
	// db := ent.NewClient(ent.Driver(drv), ent.SQLDB(drv.DB()))

	db, err := ent.Open("postgres", connectionStr)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	return db
}
