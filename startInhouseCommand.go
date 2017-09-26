package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

type inhouseIdentifier struct {
	authorChannelId string;
}

type inhouseFlow struct {
	isActive bool
}

var (
	activeFlows = make(map[inhouseIdentifier]inhouseFlow)
)

func startInhouseCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	member := getMember(s, m)
	currentGuild := getGuild(s, m)

	if memberHasRole(member, findRoleID("admin", currentGuild)) {
		authorChannel, _ := s.UserChannelCreate(m.Author.ID)

		value, exist := activeFlows[inhouseIdentifier{authorChannel.ID}]

		if exist && value.isActive {
			s.ChannelMessageSendEmbed(authorChannel.ID, &discordgo.MessageEmbed{
				Title:       "Error",
				Description: fmt.Sprintf("You tried starting an in-house in <#%s> but you have not finished setuping the previous one!", m.ChannelID) })
		} else {
			activeFlows[inhouseIdentifier{authorChannel.ID}] = inhouseFlow{true}

			s.ChannelMessageSendEmbed(authorChannel.ID,  &discordgo.MessageEmbed{
				Title:       "Starting In-House",
				Description: "Potato potatoes"})
		}
	} else {
		s.ChannelMessageSendEmbed("197502695953661952",  &discordgo.MessageEmbed{
			Title:       "Who are you",
			Description: "Potatoes potato"})
	}
}
