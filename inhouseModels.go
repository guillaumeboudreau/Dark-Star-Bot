package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

type inhouseConfig struct {
	ActiveQuestions map[string]inhouseQuestion `json:"ActiveQuestions"`
	SetupInhouses   map[string]inHouse         `json:"SetupInhouses"`
	ActiveInhouse   map[string]inHouse         `json:"ActiveInhouse"`
	InitialQuestion inhouseQuestion            `json:"InitialQuestion"`
	RolesImportance map[string]int             `json:"RolesImportance"`
}

type answerType struct {
	IsRestrictedAnswer             bool         `json:"isRestrictedAnswer"`
	ExpectedType                   reflect.Kind `json:"ExpectedType"`
	ExpectedAnswers                []string     `json:"ExpectedAnswers"`
	IsNextQuestionRelativeToAnswer bool         `json:"IsNextQuestionRelativeToAnswer"`
	InhousePropertyName            string       `json:"InhousePropertyName"`
}

type inhouseQuestion struct {
	Question               string            `json:"Question"`
	ExpectedAnswer         answerType        `json:"ExpectedAnswer"`
	NextQuestionsPerAnswer []inhouseQuestion `json:"NextQuestionsPerAnswer"`
	IsFinalQuestion        bool              `json:"IsFinalQuestion"`
	FinalMessage           string            `json:"FinalMessage"`
}

func (q inhouseQuestion) formatQuestion(channelID string) string { // Replaces the variables in the in-house flow questions by the values
	return strings.Replace(q.Question, "{inhouseChannel}", fmt.Sprintf("<#%s>", channelID), -1)
}

func (q inhouseQuestion) formatFinalMessage(channelID string) string {
	return strings.Replace(q.FinalMessage, "{inhouseChannel}", fmt.Sprintf("<#%s>", channelID), -1)
}

type inHouse struct {
	Active                     bool                 `json:"Active"`
	Name                       string               `json:"Name"`
	Date                       string               `json:"Date"`
	ChannelID                  string               `json:"ChannelID"`
	MaximumNumberOfParticipant int                  `json:"MaximumNumberOfParticipant"`
	Participants               []inHouseParticipant `json:"Participants"`
}

func (i inHouse) printInHouse() {
	log.Printf("Active : %t, Name : %s, Date : %s, ChannelID : %s", i.Active, i.Name, i.Date, i.ChannelID)
}

type inHouseParticipant struct {
	UserID     string `json:"UserID"`
	InGameName string `json:"InGameName"`
}
