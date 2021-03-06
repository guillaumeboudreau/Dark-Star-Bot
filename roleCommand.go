package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/bwmarrin/discordgo"
)

func roleCommand(s *discordgo.Session, m *discordgo.MessageCreate) { // Add role to someone
	if len(splitMsgLowered) > 1 { // If it just isnt `$role`

		for i := 0; i < len(splitMsgLowered)-1; i++ {
			assignRole(s, m, splitMsgLowered[i+1])
		}
	} else {
		s.ChannelMessageSend(m.ChannelID, "Please type a role name after `$role` !")
	}
}

func assignRole(s *discordgo.Session, m *discordgo.MessageCreate, givenRole string) {
	var roles []string
	var calls [][]string
	var serverID string
	_ = serverID // Added because it thinks it isn't being used.
	dsrFiles := getFilesFromDir("roles/*.dsr")
	for i := 0; i < len(dsrFiles); i++ {
		tcalls, troles, serverID := handledsr(dsrFiles[i])
		channel, err := s.Channel(m.ChannelID)
		if err != nil {
			log.Fatal(err)
		} else {
			guildID := channel.GuildID
			if serverID == guildID { // Is the current guild = to the current file being looped through?
				roles = troles // Then copy the roles from said file!
				calls = tcalls // And the calls!
			}
		}
	}

	roleUsed := false

	// Handles the ingame roles
	for i := 0; i < len(calls); i++ {
		for j := 0; j < len(calls[i]); j++ {
			switch givenRole {

			case calls[i][j]:
				roleUsed = true
				giveRole(s, m, roles[i])
			}
		}
	}

	if !roleUsed {
		s.ChannelMessageSend(m.ChannelID, "You cannot add that role! Use `$Roles` to see all available roles")
	}
}

func giveRole(s *discordgo.Session, m *discordgo.MessageCreate, roleNeeded string) { // Assigns the role based off of a role needed

	currentGuild := getGuild(s, m)
	currentMember := getMember(s, m)

	if currentGuild != nil && currentMember != nil {

		tempRoleID := "" // The temporary storage for role id.

		tempRoleID = findRoleID(roleNeeded, currentGuild)

		hasRole := memberHasRole(currentMember, tempRoleID)

		if !hasRole { // Give that guy a role
			err := s.GuildMemberRoleAdd(currentGuild.ID, m.Author.ID, tempRoleID) // Give the role
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "Unable to add role! Message <@!121105861539135490> and tell him there's a problem.")
				log.Fatal(err)
			}
			if err == nil { // Didnt want this popping up if there was an error
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s role added!", roleNeeded))
			}
		} else { // Bye bye role
			err := s.GuildMemberRoleRemove(currentGuild.ID, m.Author.ID, tempRoleID) // Remove the role
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "Unable to remove role! Message <@!121105861539135490> and tell him there's a problem.")
				log.Fatal(err)
			}
			if err == nil { // Didnt want this popping up if there was an error
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s role removed!", roleNeeded))
			}
		}

	}

}

func getFilesFromDir(s string) []string { // Gets all files from a directory
	r, err := filepath.Glob(s)
	if err != nil {
		log.Fatal(err)
	}
	return r
}
