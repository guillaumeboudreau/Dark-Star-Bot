package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func inhouseUnjoinCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	member := getMember(s, m)

	inhouse, inhouseExists := inhouseState.ActiveInhouse[m.ChannelID]

	if inhouseExists {
		if !userAlreadyRegistered(inhouse, member.User.ID) {
			s.ChannelMessageSendEmbed(inhouse.ChannelID, &discordgo.MessageEmbed{
				Title:       "You are not in this in house!",
				Description: fmt.Sprintf("%s, if you wish to join please use the $join command!", member.User.Mention()),
			})
		} else {
			removeUser(&inhouse, member.User.ID)
			inhouseState.ActiveInhouse[m.ChannelID] = inhouse
			saveInhouseState()
			s.ChannelMessageSendEmbed(inhouse.ChannelID, &discordgo.MessageEmbed{
				Title:       "Successfully unjoined!",
				Description: fmt.Sprintf("You have been removed from the active in-house %s!", member.User.Mention()),
			})

		}
	} else {
		s.ChannelMessageSendEmbed(inhouse.ChannelID, &discordgo.MessageEmbed{
			Title:       "In-house information!",
			Description: "There is currently no in-house active in this channel!",
		})
	}
}

func removeUser(inhouse *inHouse, userID string) {
	index := -1
	for i := 0; i < len(inhouse.Participants); i++ {
		if inhouse.Participants[i].UserID == userID {
			index = i
		}
	}

	if index != -1 {
		inhouse.Participants = append(inhouse.Participants[:index], inhouse.Participants[index+1:]...)
	}
}
