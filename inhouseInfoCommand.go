package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func inhouseInfoCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	inhouse, inhouseExists := inhouseState.ActiveInhouse[m.ChannelID]

	if inhouseExists {
		s.ChannelMessageSendEmbed(inhouse.ChannelID, &discordgo.MessageEmbed{
			Title:       "In-house information!",
			Description: "An in-house is active in this channel, use $join if you want to join the in-house!",
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:  "Name",
					Value: inhouse.Name,
				},
				&discordgo.MessageEmbedField{
					Name:  "Date",
					Value: inhouse.Date,
				},
				&discordgo.MessageEmbedField{
					Name:  "Current Number of Players",
					Value: fmt.Sprintf("%d/%d", len(inhouse.Participants), inhouse.MaximumNumberOfParticipant),
				},
			}})
	} else {
		s.ChannelMessageSendEmbed(inhouse.ChannelID, &discordgo.MessageEmbed{
			Title:       "In-house information!",
			Description: "There is currently no in-house active in this channel!",
		})
	}

}
