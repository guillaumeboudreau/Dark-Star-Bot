package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func banMessage(s *discordgo.Session, e *discordgo.GuildBanAdd) {
	cfg := loadConfig("roles/config.json")
	tserver := loadServer(cfg, e.GuildID)

	s.ChannelMessageSendEmbed(tserver.BanMessage.ID, &discordgo.MessageEmbed{
		Title:       tserver.BanMessage.Type,
		Description: fmt.Sprintf(tserver.BanMessage.Message, e.User.Mention())})
}
