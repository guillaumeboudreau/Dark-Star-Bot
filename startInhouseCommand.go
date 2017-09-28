package main

import (
	"fmt"
	"reflect"

	"github.com/bwmarrin/discordgo"
)

type inhouseIdentifier struct {
	authorChannelID string
}

type inhouseFlow struct {
	currentlyAskedQuestion inhouseQuestion
}

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

type inHouse struct {
	active    bool
	name      string
	date      string
	channelID string
}

var (
	activeFlows     = make(map[inhouseIdentifier]inhouseFlow)
	activeInhouses  = make(map[inhouseIdentifier]inHouse)
	initialQuestion = inhouseQuestion{
		"You want to start an in-house in {inhouseChannel}, what name would you want the in-house to have? (Type \"Cancel\" at anytime to cancel the in-house creation)",
		answerType{
			false,
			reflect.String,
			[]string{},
			false,
			"name",
		},
		[]inhouseQuestion{
			inhouseQuestion{
				"When is this in-house gonna take place?",
				answerType{
					false,
					reflect.String,
					[]string{},
					false,
					"date",
				},
				[]inhouseQuestion{},
				true,
				"Thank you for creating this in-house! Head over at {inhouseChannel} to see your newly create in-house!",
			},
		},
		false,
		"",
	}
)

func startInhouseCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	member := getMember(s, m)
	currentGuild := getGuild(s, m)

	if memberHasRole(member, findRoleID("admin", currentGuild)) {
		authorChannel, _ := s.UserChannelCreate(m.Author.ID)

		_, exist := activeFlows[inhouseIdentifier{authorChannel.ID}]

		if exist {
			s.ChannelMessageSendEmbed(authorChannel.ID, &discordgo.MessageEmbed{
				Title:       "Error",
				Description: fmt.Sprintf("You tried starting an in-house in <#%s> but you have not finished setuping the previous one!", m.ChannelID)})
		} else {
			activeFlows[inhouseIdentifier{authorChannel.ID}] = inhouseFlow{initialQuestion}
			activeInhouses[inhouseIdentifier{authorChannel.ID}] = inHouse{false, "", "", m.ChannelID}

			s.ChannelMessageSendEmbed(authorChannel.ID, &discordgo.MessageEmbed{
				Title:       "Starting In-House",
				Description: formatQuestion(activeFlows[inhouseIdentifier{authorChannel.ID}].currentlyAskedQuestion.question, m.ChannelID)})
		}
	} else {
		s.ChannelMessageSendEmbed("197502695953661952", &discordgo.MessageEmbed{
			Title:       "Who are you",
			Description: "Potatoes potato"})
	}
}
