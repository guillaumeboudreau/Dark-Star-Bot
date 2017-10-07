package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func welcomeMessage(s *discordgo.Session, e *discordgo.GuildMemberAdd) {
	cfg := loadConfig("roles/config.json")
	tserver := loadServer(cfg, e.GuildID)

	s.ChannelMessageSendEmbed(tserver.WelcomeMessage.ID, &discordgo.MessageEmbed{
		Title:       tserver.WelcomeMessage.Type,
		Description: fmt.Sprintf(tserver.WelcomeMessage.Message, e.User.Mention())})
}
