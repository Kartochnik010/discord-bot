package config

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	BotToken     string
	GptToken     string
	GptMaxTokens int
	GptModel     string
}

func New() (*Config, error) {
	// fetch configs from env
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		return nil, errors.New("empty token")
	}
	gptToken := os.Getenv("GPT_TOKEN")
	if gptToken == "" {
		return nil, errors.New("unset token")
	}

	gptMaxTokens, err := strconv.Atoi(os.Getenv("GPT_MAX_TOKENS"))
	if err != nil {
		return nil, errors.Join(errors.New("error while parsing gptMaxTokens: "), err)
	}
	// set default value
	if gptMaxTokens == 0 {
		gptMaxTokens = 256
	}
	gptModel := os.Getenv("GPT_MODEL")
	if gptModel == "" {
		return nil, errors.New("unset url")
	}

	// return non-empty config
	return &Config{
		BotToken:     botToken,
		GptToken:     gptToken,
		GptMaxTokens: gptMaxTokens,
		GptModel:     gptModel,
	}, nil
}
