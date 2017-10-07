package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func handleInhouseCreationFlow(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, exist := inhouseState.ActiveQuestions[m.ChannelID]

	if exist {
		if strings.TrimSpace(strings.ToLower(m.Content)) == "cancel" {
			s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
				Title:       "In-House Cancelled",
				Description: "You have cancelled the creation processus of this in-house!"})

			delete(inhouseState.ActiveQuestions, m.ChannelID)
			saveInhouseState()
		} else {
			parseInhouseCommand(s, m, m.ChannelID)
			saveInhouseState()
		}
	}
}

func parseInhouseCommand(s *discordgo.Session, m *discordgo.MessageCreate, inhouseID string) {
	activeQuestion, _ := inhouseState.ActiveQuestions[inhouseID]
	messageContentLowered := strings.TrimSpace(strings.ToLower(m.Content))
	inHouse, _ := inhouseState.SetupInhouses[inhouseID]
	expectedAnswer := activeQuestion.ExpectedAnswer

	if expectedAnswer.IsRestrictedAnswer {
		answerPosition := -1
		_ = answerPosition
		answersList := expectedAnswer.ExpectedAnswers

		for i := 0; i < len(answersList); i++ {
			stringIndex := strconv.Itoa(i)
			if strings.ToLower(answersList[0]) == messageContentLowered || stringIndex == messageContentLowered {
				answerPosition = i
				break
			}
		}

		if answerPosition == -1 {
			errorDuringQuestion(s, m.ChannelID, fmt.Errorf("Answer \"%s\" is not valid", messageContentLowered))
		} else {
			err := updateInhouseProperty(&inHouse, expectedAnswer.InhousePropertyName, expectedAnswer.ExpectedType, answersList[answerPosition])
			inhouseState.SetupInhouses[inhouseID] = inHouse
			if err == nil {
				askNextQuestion(s, m.ChannelID, inhouseID, answerPosition)
			} else {
				errorDuringQuestion(s, m.ChannelID, err)
			}
		}
	} else {
		err := updateInhouseProperty(&inHouse, expectedAnswer.InhousePropertyName, expectedAnswer.ExpectedType, m.Content)
		inhouseState.SetupInhouses[inhouseID] = inHouse
		if err == nil {
			askNextQuestion(s, m.ChannelID, inhouseID, -1)
		} else {
			errorDuringQuestion(s, m.ChannelID, err)
		}
	}
}

func askNextQuestion(s *discordgo.Session, channelID string, inhouseID string, previousAnswerPosition int) {
	activeQuestion, _ := inhouseState.ActiveQuestions[inhouseID]

	if activeQuestion.IsFinalQuestion {
		inhouse, _ := inhouseState.SetupInhouses[inhouseID]
		sendFinalMessageToUser(s, activeQuestion, channelID, inhouse.ChannelID)
		inhouse.Active = true
		inhouseState.ActiveInhouse[inhouse.ChannelID] = inhouse
		notifyInhouseActive(s, &inhouse)
		delete(inhouseState.SetupInhouses, inhouseID)
		delete(inhouseState.ActiveQuestions, channelID)
		return
	}

	if previousAnswerPosition == -1 {
		previousAnswerPosition = 0
	}

	nextQuestion := activeQuestion.NextQuestionsPerAnswer[previousAnswerPosition]
	sendNextQuestionMessage(s, nextQuestion, channelID, inhouseState.SetupInhouses[inhouseID].ChannelID)
	inhouseState.ActiveQuestions[inhouseID] = nextQuestion
}

func sendNextQuestionMessage(s *discordgo.Session, question inhouseQuestion, channelID string, inhouseChannelID string) {
	s.ChannelMessageSendEmbed(channelID, &discordgo.MessageEmbed{
		Title:       "Next question!",
		Description: question.formatQuestion(inhouseChannelID),
	})
}

func notifyInhouseActive(s *discordgo.Session, inhouse *inHouse) {
	s.ChannelMessageSendEmbed(inhouse.ChannelID, &discordgo.MessageEmbed{
		Title:       "In-house active!",
		Description: "@everyone An in-house is now active in this channel, use $join if you want to join the in-house!",
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
				Name:  "Max number of player",
				Value: strconv.Itoa(inhouse.MaximumNumberOfParticipant),
			},
		}})
}

func sendFinalMessageToUser(s *discordgo.Session, question inhouseQuestion, channelID string, inhouseID string) {
	s.ChannelMessageSendEmbed(channelID, &discordgo.MessageEmbed{
		Title:       "In-house created successfully!",
		Description: question.formatFinalMessage(inhouseID)})
}

func errorDuringQuestion(s *discordgo.Session, channelID string, err error) {
	s.ChannelMessageSendEmbed(channelID, &discordgo.MessageEmbed{
		Title:       "Error with previous answer",
		Description: fmt.Sprintf("%s, please try again!", err)})
}

func updateInhouseProperty(inHouse *inHouse, propName string, expectedType reflect.Kind, value string) error {
	ps := reflect.ValueOf(inHouse)

	s := ps.Elem()
	f := s.FieldByName(propName)
	if f.IsValid() {
		if f.CanSet() {
			if f.Kind() == expectedType {
				setInhouseValue(f, expectedType, value)
			} else {
				return fmt.Errorf("Expected property \"%s\" to be of type \"%s\"", propName, expectedType.String())
			}
		} else {
			return fmt.Errorf("Property \"%s\" cannot be set", propName)
		}
	} else {
		return fmt.Errorf("Property \"%s\" is not valid", propName)
	}
	return nil
}

func setInhouseValue(v reflect.Value, t reflect.Kind, value string) {
	switch t {
	case reflect.String:
		v.SetString(value)
		break
	}
}
