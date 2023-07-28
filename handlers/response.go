package handlers

import (
	"context"
	"errors"

	"github.com/eesoymilk/health-statistic-api/ent"
	"github.com/eesoymilk/health-statistic-api/types"
	"github.com/google/uuid"
)

func (h *Handler) RespondQuestionnaire(
	ctx context.Context,
	user_id string,
	questionnaireId uuid.UUID,
	answers []types.BaseAnswer,
) (*ent.QuestionnaireResponse, error) {
	responseNode, err := h.DB.QuestionnaireResponse.
		Create().
		SetUserID(user_id).
		SetQuestionnaireID(questionnaireId).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	for _, answer := range answers {
		answerCreate := h.DB.Answer.
			Create().
			SetQuestionID(answer.QuestionId).
			SetQuestionnaireResponse(responseNode)

		if answer.ChoiceIds != nil {
			answerCreate = answerCreate.AddChosenIDs(*answer.ChoiceIds...)
		} else if answer.Body != nil {
			answerCreate = answerCreate.SetBody(*answer.Body)
		} else {
			return nil, errors.New("both choice_ids and body in answer are nil")
		}

		_, err = answerCreate.Save(ctx)

		if err != nil {
			return nil, err
		}
	}

	return responseNode, nil
}
