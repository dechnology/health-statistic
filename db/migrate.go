package db

import (
	"context"
	"fmt"

	"github.com/eesoymilk/health-statistic-api/ent"
	"github.com/eesoymilk/health-statistic-api/ent/migrate"
)

func Migrate(ctx context.Context, db *ent.Client) error {
	// Create resources
	if err := db.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		return fmt.Errorf(
			"failed creating schema resources: %v",
			err,
		)
	}

	// Create registration questionnaire
	if err := CreateRegistrationQuestionnaire(ctx, db); err != nil {
		return fmt.Errorf(
			"failed to create registration questionnaire: %v",
			err,
		)
	}

	// Create MyCards (they are fake lol)
	if err := CreateMyCards(ctx, db); err != nil {
		return fmt.Errorf(
			"failed to create MyCards: %v",
			err,
		)
	}

	return nil
}
