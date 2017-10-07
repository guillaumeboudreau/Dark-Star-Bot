package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

var (
	inHouseConfigPath = "inhouseconfig.json"
	rolesPriorities   = map[string]int{
		"bronze":     0,
		"silver":     1,
		"gold":       2,
		"platinum":   3,
		"diamond":    4,
		"master":     5,
		"challenger": 6,
	}

	inhouseState = inhouseConfig{
		make(map[string]inhouseQuestion),
		make(map[string]inHouse),
		make(map[string]inHouse),
		inhouseQuestion{
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
					[]inhouseQuestion{
						inhouseQuestion{
							"How many players can take part in this in-house? (Normally 10)",
							answerType{
								false,
								reflect.Int,
								[]string{},
								false,
								"MaximumNumberOfParticipant",
							},
							[]inhouseQuestion{},
							true,
							"Thank you for creating this in-house! Head over at {inhouseChannel} to see your newly created in-house!",
						},
					},
					false,
					"",
				},
			},
			false,
			"",
		},
		rolesPriorities,
	}
)

func saveInhouseState() {
	currentStateJSON, err := json.MarshalIndent(inhouseState, "", "    ")
	if err != nil {
		fmt.Println("Error during json conversion : ", err.Error())
	}
	ioutil.WriteFile(inHouseConfigPath, currentStateJSON, 0644)
}

func loadInhouseState() {
	raw, err := ioutil.ReadFile(inHouseConfigPath)
	if err != nil {
		fmt.Println("json file not found!, ", err.Error())
	}
	json.Unmarshal(raw, &inhouseState)
}
