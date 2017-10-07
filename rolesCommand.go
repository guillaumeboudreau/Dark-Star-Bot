package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func rolesCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	cfg := loadConfig("roles/config.json")
	channel, err := s.Channel(m.ChannelID)
	_ = channel
	if err != nil {
		fmt.Print("Error getting current channel: ", err)
	}

	tserver := loadServer(cfg, channel.GuildID)

	var ps string
	if len(tserver.Roles) > 0 {
		for i := 0; i < len(tserver.Roles); i++ {
			ps += fmt.Sprintf("%s\n", tserver.Roles[i].Role)
		}
	} else {
		ps = "No roles available!"
	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{ // Straight up prints all the roles you can have.. Nothing else to it
		Title:       "Roles Available:",
		Description: ps})
}
