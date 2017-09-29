package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

type answerType struct {
	isRestrictedAnswer             bool
	expectedType                   reflect.Kind
	expectedAnswers                []string
	isNextQuestionRelativeToAnswer bool
	inhousePropertyName            string
}

type inhouseQuestion struct {
	question               string
	expectedAnswer         answerType
	nextQuestionsPerAnswer []inhouseQuestion
	isFinalQuestion        bool
	finalMessage           string
}

func (q inhouseQuestion) formatQuestion(channelID string) string { // Replaces the variables in the in-house flow questions by the values
	return strings.Replace(q.question, "{inhouseChannel}", fmt.Sprintf("<#%s>", channelID), -1)
}

func (q inhouseQuestion) formatFinalMessage(channelID string) string {
	return strings.Replace(q.finalMessage, "{inhouseChannel}", fmt.Sprintf("<#%s>", channelID), -1)
}

type inHouse struct {
	Active    bool
	Name      string
	Date      string
	ChannelID string
}

func (i inHouse) printInHouse() {
	log.Printf("Active : %t, Name : %s, Date : %s, ChannelID : %s", i.Active, i.Name, i.Date, i.ChannelID)
}
