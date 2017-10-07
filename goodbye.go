package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func goodbyeMessage(s *discordgo.Session, e *discordgo.GuildMemberRemove) {
	cfg := loadConfig("roles/config.json")
	tserver := loadServer(cfg, e.GuildID)

	s.ChannelMessageSendEmbed(tserver.GoodbyeMessage.ID, &discordgo.MessageEmbed{
		Title:       tserver.GoodbyeMessage.Type,
		Description: fmt.Sprintf(tserver.GoodbyeMessage.Message, e.User.Mention())})
}
