package main

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func handleInhouseCreationFlow(s *discordgo.Session, m *discordgo.MessageCreate) {
	inHouseID := inhouseIdentifier{m.ChannelID}
	_, exist := activeFlows[inHouseID]

	if exist {
		if m.Content == "cancel" {
			s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
				Title:       "In-House Cancelled",
				Description: "You have cancelled the creation processus of this in-house!"})

			delete(activeFlows, inHouseID)
		} else {
			parseInhouseCommand(s, m, inHouseID)
		}
	}
}

func parseInhouseCommand(s *discordgo.Session, m *discordgo.MessageCreate, inhouseID inhouseIdentifier) {
	inhouseFlow, _ := activeFlows[inhouseID]
	messageContentLowered := strings.TrimSpace(strings.ToLower(m.Content))
	inHouse, _ := activeInhouses[inhouseID]
	expectedAnswer := inhouseFlow.currentlyAskedQuestion.expectedAnswer

	if expectedAnswer.isRestrictedAnswer {
		answerPosition := -1
		_ = answerPosition
		answersList := expectedAnswer.expectedAnswers

		for i := 0; i < len(answersList); i++ {
			stringIndex := strconv.Itoa(i)
			if strings.ToLower(answersList[0]) == messageContentLowered || stringIndex == messageContentLowered {
				answerPosition = i
				break
			}
		}

		if answerPosition == -1 {
			// Handle ask question again
		} else {
			updateInhouse(inHouse, expectedAnswer.inhousePropertyName, expectedAnswer.expectedType, answersList[answerPosition])
		}
	}
}

func updateInhouse(inHouse inHouse, propName string, expectedType reflect.Kind, value string) {
	ps := reflect.ValueOf(&inHouse)

	s := ps.Elem()
	if s.Kind() == reflect.Struct {
		f := s.FieldByName(propName)
		if f.IsValid() {
			if f.CanSet() {
				if f.Kind() == expectedType {
					setValue(f, expectedType, value)
				}
			}
		}
	}
}

func setValue(v reflect.Value, t reflect.Kind, value string) {
	switch t {
	case reflect.String:
		v.SetString(value)
		break
	}
}
