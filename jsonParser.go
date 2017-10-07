package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

type dsgRole struct { // To hold calls and the role those calls affect
	Calls []string `json:"calls"`
	Role  string   `json:"role"`
}

type dsgMessage struct { // To hold a message and a channel to send when using dsgMessage.sendMessage()
	ID      string `json:"ID"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

type dsgServer struct {
	ServerID       string     `json:"serverID"`
	Roles          []dsgRole  `json:"roles"`
	WelcomeMessage dsgMessage `json:"welcomeMessage"`
	GoodbyeMessage dsgMessage `json:"goodbyeMessage"`
	BanMessage     dsgMessage `json:"banMessage"`
}

type dsgConfig struct {
	Files []dsgFile `json:"files"`
}

type dsgFile struct {
	ID       string `json:"ID"`
	Location string `json:"location"`
}

func loadConfig(l string) dsgConfig {
	var cfg dsgConfig
	raw, err := ioutil.ReadFile(l)
	if err != nil {
		fmt.Println("json file not found!, ", err.Error())
		var r dsgConfig
		return r
	}
	json.Unmarshal(raw, &cfg)
	return cfg
}

func loadServers(c dsgConfig) []dsgServer {
	servers := make([]dsgServer, len(c.Files))
	for i := 0; i < len(c.Files); i++ {
		var s dsgServer
		raw, err := ioutil.ReadFile(c.Files[i].Location)
		if err != nil {
			fmt.Println("json file not found!, ", err.Error())
			var r []dsgServer
			return r
		}
		json.Unmarshal(raw, &s)
		servers = append(servers, s)
	}

	return servers
}

func loadServer(c dsgConfig, ID string) dsgServer {
	i := 0
	for i = 0; i < len(c.Files); i++ {
		if c.Files[i].ID == ID {
			var s dsgServer
			raw, err := ioutil.ReadFile(c.Files[i].Location)
			if err != nil {
				fmt.Println("json file not found!, ", err.Error())
				var r dsgServer
				return r
			}
			err = json.Unmarshal(raw, &s)
			if err != nil {
				fmt.Println("Error unmarshal-ing JSON file, ", err.Error())
				var r dsgServer
				return r
			}

			return s
		}
	}
	var r dsgServer
	return r
}

//SpaceMap ... removes all whitespace from a string efficiently
func SpaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}
