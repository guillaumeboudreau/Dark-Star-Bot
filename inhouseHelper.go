package main

import (
	"fmt"
	"strings"
)

func formatQuestion(s string, channelID string) string { // Replaces the variables in the in-house flow questions by the values
	return strings.Replace(s, "{channelId}", fmt.Sprintf("<#%s>", channelID), -1)
}
