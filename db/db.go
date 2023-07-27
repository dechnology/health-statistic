package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/eesoymilk/health-statistic-api/ent"
	"github.com/eesoymilk/health-statistic-api/ent/migrate"
	"github.com/eesoymilk/health-statistic-api/handlers"
	"github.com/eesoymilk/health-statistic-api/types"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func ReadRegistrationQuestionnaire() (*types.QuestionnaireWithId, error) {
	data, err := os.ReadFile("./db/registration_questionnaire.json")
	if err != nil {
		return nil, err
	}

	// Parse the JSON
	var questionnaireData types.QuestionnaireWithId
	err = json.Unmarshal(data, &questionnaireData)
	if err != nil {
		return nil, err
	}

	// Print the json
	jsonBytes, err := json.MarshalIndent(questionnaireData, "", "  ")
	if err != nil {
		return nil, err
	}
	log.Printf("regisration questionnaire: %v", string(jsonBytes))

	return &questionnaireData, nil
}

func CreateRegistrationQuestionnaire(
	ctx context.Context,
	db *ent.Client,
) error {
	// Parse the JSON
	questionnaireData, err := ReadRegistrationQuestionnaire()
	if err != nil {
		return fmt.Errorf("error marshaling to JSON: %v", err)
	}

	id, err := uuid.Parse(questionnaireData.ID)
	if err != nil {
		return fmt.Errorf("failed parsing id: %v", err)
	}

	_, err = db.Questionnaire.Create().
		SetID(id).
		SetName(questionnaireData.Name).
		Save(ctx)
	if err != nil {
		// This error occurs if registration questionnaire is already created
		return fmt.Errorf(
			"failed to create registration questionnaire: %v",
			err,
		)
	}

	h := handlers.Handler{DB: db}
	if err := h.AppendQuestions(
		ctx,
		questionnaireData.ID,
		questionnaireData.Questions,
	); err != nil {
		return fmt.Errorf(
			"failed to append registration questionnaire questions: %v",
			err,
		)
	}

	return nil
}

func ReadMyCards() (*[]types.BaseMyCard, error) {
	data, err := os.ReadFile("./db/mycards.json")
	if err != nil {
		return nil, err
	}

	// Parse the JSON
	var mycards []types.BaseMyCard
	err = json.Unmarshal(data, &mycards)
	if err != nil {
		return nil, err
	}

	// Print the json
	jsonBytes, err := json.MarshalIndent(mycards, "", "  ")
	if err != nil {
		return nil, err
	}
	log.Printf("fake mycards: %v", string(jsonBytes))

	return &mycards, nil
}

func CreateMyCards(ctx context.Context, db *ent.Client) error {
	// Parse the JSON
	mycards, err := ReadMyCards()
	if err != nil {
		return fmt.Errorf("error marshaling to JSON: %v", err)
	}

	for _, mycardData := range *mycards {
		_, err = db.MyCard.Create().
			SetID(mycardData.CardNumber).
			SetCardPassword(mycardData.CardPassword).
			Save(ctx)
		if err != nil {
			// This error occurs if registration questionnaire is already created
			return fmt.Errorf(
				"failed to create mycard: %v",
				err,
			)
		}
	}

	return nil
}

func Migrate(ctx context.Context, db *ent.Client) error {
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

	if err := CreateRegistrationQuestionnaire(ctx, db); err != nil {
		return fmt.Errorf(
			"failed to create registration questionnaire: %v",
			err,
		)
	}

	if err := CreateMyCards(ctx, db); err != nil {
		return fmt.Errorf(
			"failed to create MyCards: %v",
			err,
		)
	}

	return nil
}

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

	db, err := ent.Open("postgres", connectionStr)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	return db
}
