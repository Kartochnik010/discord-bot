package handler

import (
	"log/slog"
	"strings"

	"github.com/Kartochnik010/discord-bot/internal/api"
	"github.com/Kartochnik010/discord-bot/internal/config"
	"github.com/Kartochnik010/discord-bot/internal/lib/logger/sl"
	"github.com/bwmarrin/discordgo"
)

const (
	prefix = "moris "
)

func Handler(cfg *config.Config) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// return if the author of the incoming message the bot itself
		if m.Author.ID == s.State.User.ID {
			return
		}

		// check if message starts with needed prefix return if not
		if strings.Index(m.Content, prefix) != 0 {
			return
		}
		// trim message content
		m.Content = m.Content[len(prefix):]

		mp := map[string]func(){
			"hello": func() {
				s.ChannelMessageSend(m.ChannelID, "Hi!")
			},
			"yt": func() {
				s.ChannelMessageSend(m.ChannelID, "https://www.youtube.com/watch?v=2GgiZZhO-PA")
			},
			"help": func() {
				s.ChannelMessageSend(m.ChannelID, "```\nhello - Bot will answer with Hi! Kind of a healthcheck\nentertain - link to random favourite yt videos\nhelp - print this message```\n\nYou can also type `moris ...` to ask me anything :)")
			},
		}

		if f, ok := mp[m.Content]; ok {
			f()
		} else {
			gptRes, err := api.PromptGPT(cfg, m.Content)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "Sorry, I am busy at the moment. Try again later.")
				slog.Error("failed to promt gpt", sl.Err(err))
				return
			}
			gptResContent := gptRes.Choices[len(gptRes.Choices)-1].Message.Content

			s.ChannelMessageSend(m.ChannelID, gptResContent)
		}
	}
}
