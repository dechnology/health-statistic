package db

import (
	"log"

	"github.com/eesoymilk/health-statistic-api/ent"
	_ "github.com/lib/pq"
)

func New() *ent.Client {
	db, err := ent.Open(
		"postgres",
		"host=localhost user=postgres dbname=health-statistic password=510600 sslmode=disable",
	)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	return db
}
