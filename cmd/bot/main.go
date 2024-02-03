package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		slog.Error("failed to init bot")
		os.Exit(1)
	}
	session, err := discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		slog.Error("failed to init bot", Err(err))
		os.Exit(1)
	}

	session.AddHandler(helloHandler)

	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = session.Open()
	if err != nil {
		slog.Error("failed to open discord session", Err(err))
		os.Exit(1)
	}
	defer session.Close()

	slog.Info("Bot")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGTERM, syscall.SIGINT)
	<-sc
	slog.Info("Stopping bot")
}

func helloHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// return if the author of the incoming message the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}
	mp := map[string]func(){
		"hello": func() {
			s.ChannelMessageSend(m.ChannelID, "hi")
		},
		"bye": func() {
			s.ChannelMessageSend(m.ChannelID, "Good bye")
		},
		"commands": func() {
			s.ChannelMessageSend(m.ChannelID, "hello\nbye")
		},
	}

	if f, ok := mp[m.Content]; ok {
		f()
	}

}

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
