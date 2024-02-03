package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Kartochnik010/discord-bot/internal/config"
	"github.com/Kartochnik010/discord-bot/internal/models"
)

var (
	client = &http.Client{}
)

func PromptGPT(cfg *config.Config, msg string) (*models.GptResponse, error) {
	gptRequest := &models.GptRequest{
		Model: "https://api.openai.com/v1/chat/completions",
		Messages: []models.Message{
			{
				Role:    "system",
				Content: "I need you to act like a moris from Madagaskar movie. answering this message you DO NOT BREAK CHARACTER!!! DO NOT TELL ME THAT YOU ARE AN AI!!!",
			},
			{
				Role:    "user",
				Content: msg,
			},
		},
		Temperature: 0.2,
		MaxTokens:   cfg.GptMaxTokens,
	}

	return SendRequest(gptRequest, cfg)
}

func SendRequest(gptRequest *models.GptRequest, cfg *config.Config) (*models.GptResponse, error) {
	b, err := json.Marshal(gptRequest)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	fmt.Println("request: ", string(b))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", strings.Trim(cfg.GptToken, "\"")))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.Status != "200 OK" {
		return nil, errors.New(resp.Status)
	}

	res := &models.GptResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
