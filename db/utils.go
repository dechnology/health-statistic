package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/eesoymilk/health-statistic-api/ent"
	"github.com/eesoymilk/health-statistic-api/handlers"
	"github.com/eesoymilk/health-statistic-api/types"
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

	_, err = db.Questionnaire.Create().
		SetID(questionnaireData.ID).
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
