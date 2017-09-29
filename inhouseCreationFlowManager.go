package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func handleInhouseCreationFlow(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, exist := activeQuestions[m.ChannelID]

	if exist {
		if strings.TrimSpace(strings.ToLower(m.Content)) == "cancel" {
			s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
				Title:       "In-House Cancelled",
				Description: "You have cancelled the creation processus of this in-house!"})

			delete(activeQuestions, m.ChannelID)
		} else {
			parseInhouseCommand(s, m, m.ChannelID)
		}
	}
}

func parseInhouseCommand(s *discordgo.Session, m *discordgo.MessageCreate, inhouseID string) {
	activeQuestion, _ := activeQuestions[inhouseID]
	messageContentLowered := strings.TrimSpace(strings.ToLower(m.Content))
	inHouse, _ := setupInhouses[inhouseID]
	// inHouse.printInHouse()
	expectedAnswer := activeQuestion.expectedAnswer

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
			errorDuringQuestion(s, m.ChannelID, fmt.Errorf("Answer \"%s\" is not valid", messageContentLowered))
		} else {
			err := updateInhouseProperty(inHouse, expectedAnswer.inhousePropertyName, expectedAnswer.expectedType, answersList[answerPosition])
			if err == nil {
				askNextQuestion(s, m.ChannelID, inhouseID, answerPosition)
			} else {
				errorDuringQuestion(s, m.ChannelID, err)
			}
		}
	} else {
		err := updateInhouseProperty(inHouse, expectedAnswer.inhousePropertyName, expectedAnswer.expectedType, messageContentLowered)
		if err == nil {
			askNextQuestion(s, m.ChannelID, inhouseID, -1)
		} else {
			errorDuringQuestion(s, m.ChannelID, err)
		}
	}
	inHouse.printInHouse()
}

func askNextQuestion(s *discordgo.Session, channelID string, inhouseID string, previousAnswerPosition int) {
	activeQuestion, _ := activeQuestions[inhouseID]

	if activeQuestion.isFinalQuestion {
		inhouse, _ := setupInhouses[inhouseID]
		sendFinalMessageToUser(s, activeQuestion, channelID, inhouse.ChannelID)
		inhouse.Active = true
		activeInhouse[inhouse.ChannelID] = inhouse
		notifyInhouseActive(s, inhouse)
		delete(setupInhouses, inhouseID)
		delete(activeQuestions, channelID)
		return
	}

	if previousAnswerPosition == -1 {
		previousAnswerPosition = 0
	}

	nextQuestion := activeQuestion.nextQuestionsPerAnswer[previousAnswerPosition]
	sendNextQuestionMessage(s, nextQuestion, channelID, setupInhouses[inhouseID].ChannelID)
	activeQuestions[inhouseID] = nextQuestion
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
	fmt.Printf("Now in updateInhouseProperty with propName : %s ExpectedType : %s and value : %s\n", propName, expectedType.String(), value)

	ps := reflect.ValueOf(inHouse)

	s := ps.Elem()
	f := s.FieldByName(propName)
	if f.IsValid() {
		if f.CanSet() {
			if f.Kind() == expectedType {
				setInhouseValue(f, expectedType, value)
			} else {
				fmt.Println("Is not of expected kind")
				return fmt.Errorf("Expected property \"%s\" to be of type \"%s\"", propName, expectedType.String())
			}
		} else {
			fmt.Println("Cannot set")
			return fmt.Errorf("Property \"%s\" cannot be set", propName)
		}
	} else {
		fmt.Println("Is not valid")
		return fmt.Errorf("Property \"%s\" is not valid", propName)
	}
	return nil
}

func setInhouseValue(v reflect.Value, t reflect.Kind, value string) {
	fmt.Println("Now in set inhouse")
	switch t {
	case reflect.String:
		v.SetString(value)
		break
	}
}
