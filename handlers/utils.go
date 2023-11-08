package handlers

import (
	"context"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/eesoymilk/health-statistic-api/ent"
	"github.com/eesoymilk/health-statistic-api/ent/user"
	"github.com/eesoymilk/health-statistic-api/middlewares"
	"github.com/gin-gonic/gin"
)

// "client_id": "Ge4HnGw9b5tKnU1V5wh0b1eYluHLVcjt"
// "client_secret": "O_e4lZMtK96qOb8emZiFEOzAEW4pHSpVoXBND1V7pnf1kh2e_e-ZjKiKMtFXtxUR"
// "audience": "https://health-statistic.dechnology.com.tw/api/v2/"
// "grant_type": "client_credentials"

func GetManagementToken() (*string, error) {
	auth0Issuer := os.Getenv("AUTH0_ISSUER_URL")
	clientId := os.Getenv("AUTH0_CLIENT_ID")
	clientSecret := os.Getenv("AUTH0_CLIENT_SECRET")

	if auth0Issuer == "" || clientId == "" || clientSecret == "" {
		return nil, nil
	}

	url := auth0Issuer + "/oauth/token"
	payload := strings.NewReader("grant_type=client_credentials&client_id=" + clientId + "&client_secret=" + clientSecret + "&audience=" + auth0Issuer + "/api/v2/")
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		return nil, err
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	managementToken := string(body)

	return &managementToken, nil
}

func GetUserId(c *gin.Context) (*string, error) {
	token, err := middlewares.GetTokenInGin(c)
	if err != nil {
		return nil, err
	}
	return &token.RegisteredClaims.Subject, nil
}

func (h *Handler) GetUserById(
	ctx context.Context, userId string,
) (*ent.User, error) {
	userNode, err := h.DB.User.Query().
		Where(user.ID(userId)).
		WithQuestionnaireResponses(func(qrq *ent.QuestionnaireResponseQuery) {
			qrq.WithQuestionnaire().WithAnswers(func(aq *ent.AnswerQuery) {
				aq.WithChosen().WithQuestion()
			})
		}).
		WithMycards().
		WithPrices().
		WithDeegoo().
		WithHealthkit().
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return userNode, nil
}
