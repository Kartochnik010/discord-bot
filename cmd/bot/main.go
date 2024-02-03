package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/Kartochnik010/discord-bot/internal/config"
	"github.com/Kartochnik010/discord-bot/internal/handler"
	"github.com/Kartochnik010/discord-bot/internal/lib/logger/sl"
	"github.com/bwmarrin/discordgo"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		slog.Error("failed to load config", sl.Err(err))
		os.Exit(1)
	}

	session, err := discordgo.New(fmt.Sprintf("Bot %s", cfg.BotToken))
	if err != nil {
		slog.Error("failed to init bot", sl.Err(err))
		os.Exit(1)
	}

	// Handler handles messages asyncrounosly from the main goroutine
	session.AddHandler(handler.Handler(cfg))

	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = session.Open()
	if err != nil {
		slog.Error("failed to open discord session", sl.Err(err))
		os.Exit(1)
	}
	defer session.Close()

	slog.Info("Bot has started")

	// Graceful shutdown
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGTERM, syscall.SIGINT)
	<-sc
	slog.Info("Stopping bot")
}
