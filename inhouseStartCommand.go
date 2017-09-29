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

		_, exist := activeQuestions[authorChannel.ID]

		if exist {
			s.ChannelMessageSendEmbed(authorChannel.ID, &discordgo.MessageEmbed{
				Title:       "Error",
				Description: fmt.Sprintf("You tried starting an in-house in <#%s> but you have not finished setuping the previous one!", m.ChannelID)})
		} else {
			activeQuestions[authorChannel.ID] = initialQuestion
			setupInhouses[authorChannel.ID] = &inHouse{false, "In-House", "N/A", m.ChannelID}
			s.ChannelMessageSendEmbed(authorChannel.ID, &discordgo.MessageEmbed{
				Title:       "Starting In-House",
				Description: activeQuestions[authorChannel.ID].formatQuestion(m.ChannelID)})
		}
	} else {
		s.ChannelMessageSendEmbed("197502695953661952", &discordgo.MessageEmbed{
			Title:       "Who are you",
			Description: "Potatoes potato"})
	}
}
