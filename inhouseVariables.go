package main

import (
	"reflect"
)

var (
	activeQuestions = make(map[string]inhouseQuestion)
	setupInhouses   = make(map[string]*inHouse)
	activeInhouse   = make(map[string]*inHouse)
	initialQuestion = inhouseQuestion{
		"You want to start an in-house in {inhouseChannel}, what name would you want the in-house to have? (Type \"Cancel\" at anytime to cancel the in-house creation)",
		answerType{
			false,
			reflect.String,
			[]string{},
			false,
			"Name",
		},
		[]inhouseQuestion{
			inhouseQuestion{
				"When is this in-house gonna take place?",
				answerType{
					false,
					reflect.String,
					[]string{},
					false,
					"Date",
				},
				[]inhouseQuestion{},
				true,
				"Thank you for creating this in-house! Head over at {inhouseChannel} to see your newly created in-house!",
			},
		},
		false,
		"",
	}
)
