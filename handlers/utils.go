package handlers

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/eesoymilk/health-statistic-api/ent"
	"github.com/eesoymilk/health-statistic-api/ent/user"
	"github.com/eesoymilk/health-statistic-api/middlewares"
	"github.com/eesoymilk/health-statistic-api/types"
	"github.com/gin-gonic/gin"
)

func GetManagementToken() (*string, error) {
	auth0Issuer := os.Getenv("AUTH0_ISSUER_URL")
	clientId := os.Getenv("AUTH0_CLIENT_ID")
	clientSecret := os.Getenv("AUTH0_CLIENT_SECRET")

	if auth0Issuer == "" || clientId == "" || clientSecret == "" {
		return nil, nil
	}

	url := auth0Issuer + "oauth/token"
	payload := strings.NewReader("grant_type=client_credentials&client_id=" + clientId + "&client_secret=" + clientSecret + "&audience=" + auth0Issuer + "api/v2/")
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		return nil, err
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// Unmarshal the JSON data into the struct
	var responseData types.ManagementTokenResponse
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return nil, err
	}

	// Now you can use responseData.AccessToken
	accessToken := responseData.AccessToken
	log.Print(accessToken)

	return &accessToken, nil
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
