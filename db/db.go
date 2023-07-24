package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/eesoymilk/health-statistic-api/ent"
	"github.com/eesoymilk/health-statistic-api/types"
	_ "github.com/lib/pq"
)

func ReadRegistrationQuestionnaire() (*types.QuestionnaireWithId, error) {
	data, err := os.ReadFile("./db/registration_questionnaire.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	// Parse the JSON
	var questionnaireData types.QuestionnaireWithId
	err = json.Unmarshal(data, &questionnaireData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	// Print the json
	jsonBytes, err := json.MarshalIndent(questionnaireData, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("error marshaling to JSON: %v", err)
	}
	log.Printf("data: %v", string(jsonBytes))

	return &questionnaireData, nil
}

func Migrate(db *ent.Client) error {
	if err := db.Schema.Create(context.Background()); err != nil {
		return fmt.Errorf("failed creating schema resources: %v", err)
	}

	// Parse the JSON
	questionnaireData, err := ReadRegistrationQuestionnaire()
	if err != nil {
		return err
	}

	questionnaireNode, err := db.Questionnaire.Create().
		SetName(questionnaireData.Name).
		Save(context.Background())

	for i, questionData := range questionnaireData.Questions {
		_, err := db.Question.Create().
			SetBody(questionData.Body).
			SetOrder(i).
			SetQuestionnaire(questionnaireNode).
			Save(context.Background())

		if err != nil {
			return fmt.Errorf("failed to create question: %v", err)
		}
	}

	if err != nil {
		return fmt.Errorf("failed to parse JSON: %v", err)
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
