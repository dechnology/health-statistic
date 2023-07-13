package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Payload struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Audience     string `json:"audience"`
	GrantType    string `json:"grant_type"`
}

type ResponseBody struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	payloadData := Payload{
		ClientID:     os.Getenv("AUTH0_CLIENT_ID"),
		ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
		Audience:     os.Getenv("AUTH0_AUDIENCE"),
		GrantType:    "client_credentials",
	}

	payloadBytes, err := json.Marshal(payloadData)
	if err != nil {
		log.Fatalf("Error marshalling the payload to JSON: %v", err)
	}

	payload := bytes.NewReader(payloadBytes)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	fmt.Printf("Status: %s\n", res.Status)
	fmt.Println("Headers:")
	for name, headers := range res.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			fmt.Printf("%v: %v\n", name, h)
		}
	}

	body, _ := io.ReadAll(res.Body)

	var responseBody ResponseBody
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		log.Fatalf("Error parsing JSON response body: %v", err)
	}

	fmt.Printf("\nBody:\nAccess Token: %s\nExpires In: %d\nToken Type: %s\n", responseBody.AccessToken, responseBody.ExpiresIn, responseBody.TokenType)

}
