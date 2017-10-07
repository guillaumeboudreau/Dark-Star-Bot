package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func startInhouseCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	member := getMember(s, m)
	currentGuild := getGuild(s, m)

	if memberHasRole(member, findRoleID("admin", currentGuild)) {
		authorChannel, _ := s.UserChannelCreate(m.Author.ID)

		_, exist := inhouseState.ActiveQuestions[authorChannel.ID]
		_, inhouseExists := inhouseState.ActiveInhouse[m.ChannelID]

		if exist {
			s.ChannelMessageSendEmbed(authorChannel.ID, &discordgo.MessageEmbed{
				Title:       "Error",
				Description: fmt.Sprintf("You tried starting an in-house in <#%s> but you have not finished setuping the previous one!", m.ChannelID)})
		} else if inhouseExists {
			s.ChannelMessageSendEmbed(authorChannel.ID, &discordgo.MessageEmbed{
				Title:       "Error",
				Description: fmt.Sprintf("You tried starting an in-house in <#%s> but there is already an existing one!", m.ChannelID)})
		} else {
			inhouseState.ActiveQuestions[authorChannel.ID] = inhouseState.InitialQuestion
			inhouseState.SetupInhouses[authorChannel.ID] = inHouse{false, "In-House", "N/A", m.ChannelID, 10, []inHouseParticipant{}}
			s.ChannelMessageSendEmbed(authorChannel.ID, &discordgo.MessageEmbed{
				Title:       "Starting In-House",
				Description: inhouseState.ActiveQuestions[authorChannel.ID].formatQuestion(m.ChannelID)})
			saveInhouseState()
		}
	}
	s.ChannelMessageDelete(m.ChannelID, m.ID)
}
