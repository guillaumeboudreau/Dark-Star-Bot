package main

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/bwmarrin/discordgo"
)

type participantPriority struct {
	ID       string
	Priority int
}

type byPriority []participantPriority

func (a byPriority) Len() int           { return len(a) }
func (a byPriority) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byPriority) Less(i, j int) bool { return a[i].Priority < a[j].Priority }

func inhouseGenerateCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	inhouse, inhouseExists := inhouseState.ActiveInhouse[m.ChannelID]

	if inhouseExists {
		if len(inhouse.Participants)%2 != 0 {
			s.ChannelMessageSendEmbed(inhouse.ChannelID, &discordgo.MessageEmbed{
				Title:       "Bracket not generated!",
				Description: "The number of participant is odd so bracket cannot be generated",
			})
		} else {
			priorities := []participantPriority{}

			for i := 0; i < len(inhouse.Participants); i++ {
				userID := inhouse.Participants[i].UserID
				priority := getPriorityForMember(s, m, userID)
				priorities = append(priorities, participantPriority{userID, priority})
			}

			team1, team2 := generateBracket(priorities)

			var buffer bytes.Buffer
			buffer.WriteString("Team 1 : \n")
			for i := 0; i < len(team1); i++ {
				buffer.WriteString(mentionUser(team1[i].ID))
				buffer.WriteString("\n\n")
			}
			buffer.WriteString("Team 2 : \n")
			for i := 0; i < len(team2); i++ {
				buffer.WriteString(mentionUser(team1[2].ID))
				buffer.WriteString("\n\n")
			}
			s.ChannelMessageSendEmbed(inhouse.ChannelID, &discordgo.MessageEmbed{
				Title:       "In-house bracket generated!",
				Description: buffer.String(),
			})

		}
	} else {
		s.ChannelMessageSendEmbed(inhouse.ChannelID, &discordgo.MessageEmbed{
			Title:       "In-house information!",
			Description: "There is currently no in-house active in this channel!",
		})
	}
}

func getPriorityForMember(s *discordgo.Session, m *discordgo.MessageCreate, userID string) int {
	currentGuild := getGuild(s, m)
	member, err := s.State.Member(currentGuild.ID, userID)
	if err != nil {
		fmt.Println("Error making state", err)
	}

	priority := 0
	for k, v := range inhouseState.RolesImportance {
		tempRoleID := findRoleID(k, currentGuild)
		if memberHasRole(member, tempRoleID) {
			priority = v
			break
		}
	}

	return priority
}

func generateBracket(participants []participantPriority) ([]participantPriority, []participantPriority) {
	sort.Sort(byPriority(participants))
	set1 := []participantPriority{}
	set2 := []participantPriority{}

	for i := 0; i < len(participants); i++ {
		set1 = append(set1, participants[i])
		i++
		set2 = append(set2, participants[i])
		tempset := set2
		set2 = set1
		set1 = tempset
	}
	return set1, set2
}

func containsParticipant(participants []participantPriority, participant participantPriority) bool {
	for i := 0; i < len(participants); i++ {
		if participants[i] == participant {
			return true
		}
	}
	return false
}

func mentionUser(userID string) string {
	return "<@" + userID + ">"
}
