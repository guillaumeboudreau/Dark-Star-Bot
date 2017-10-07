package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func inhouseJoinCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	member := getMember(s, m)

	inhouse, inhouseExists := inhouseState.ActiveInhouse[m.ChannelID]

	if inhouseExists {
		if userAlreadyRegistered(inhouse, member.User.ID) {
			s.ChannelMessageSendEmbed(inhouse.ChannelID, &discordgo.MessageEmbed{
				Title:       "You have already joined this in house",
				Description: fmt.Sprintf("%s, if you wish to unjoin please use the $unjoin command!", member.User.Mention()),
			})
		} else {
			inhouse.Participants = append(inhouse.Participants, inHouseParticipant{member.User.ID, ""})
			inhouseState.ActiveInhouse[m.ChannelID] = inhouse
			saveInhouseState()
			if len(inhouse.Participants) >= inhouse.MaximumNumberOfParticipant {
				s.ChannelMessageSendEmbed(inhouse.ChannelID, &discordgo.MessageEmbed{
					Title:       "Thank you for joining",
					Description: fmt.Sprintf("The in-house was already full but you have been added as a sub %s!", member.User.Mention()),
				})
			} else {
				s.ChannelMessageSendEmbed(inhouse.ChannelID, &discordgo.MessageEmbed{
					Title:       "Thank you for joining!",
					Description: fmt.Sprintf("You have been added in the active in-house %s!", member.User.Mention()),
				})
			}
		}
	} else {
		s.ChannelMessageSendEmbed(inhouse.ChannelID, &discordgo.MessageEmbed{
			Title:       "In-house information!",
			Description: "There is currently no in-house active in this channel!",
		})
	}
}

func userAlreadyRegistered(inhouse inHouse, userID string) bool {
	for i := 0; i < len(inhouse.Participants); i++ {
		if inhouse.Participants[i].UserID == userID {
			return true
		}
	}
	return false
}
