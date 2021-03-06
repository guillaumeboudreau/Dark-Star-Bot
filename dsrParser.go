package main

import (
	"log"
	"os"
	"strings"
	"unicode"
)

/*
 * WHEN MAKING THE .DSR FILES, YOU MUST FOLLOW THIS FORMAT:
 * LINE 1: Always the server ID with a ; on the end
 *
 * LINES PROCEEDING must be done like so:
 * CALL,CALL,CALL=ROLE;
 *
 * Calls cannot have spaces due to the fact you can call
 * multiple roles with a single command, ex: $role a b c
 *
 * But; the role itself CAN have spaces. An example of a line would be:
 * role,arole=A Role;
 */

func handledsr(filename string) ([][]string, []string, string) { // Opens a dsr file and returns the role calls,
	//													and then the actual role
	file, err := os.Open(filename) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 5000)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	s := string(data[:count])
	fLines := strings.Split(s, ";") // Separates lines
	// Removed the first line, which is the server ID
	sID := fLines[0]
	fLines = fLines[1:]
	var fCalls []string
	var fRoles []string
	for i := 0; i < len(fLines); i++ {
		t := strings.Split(fLines[i], "=")
		t[0] = SpaceMap(t[0])
		if len(t) >= 2 {
			fCalls = append(fCalls, t[0]) // Should get the information before the '=' on that line
			fRoles = append(fRoles, t[1]) // Should get the information after the '=' on that line
		}
	}

	var fRolesFinal []string
	var fCallsFinal [][]string
	for i := 0; i < len(fRoles); i++ {
		fRolesFinal = append(fRolesFinal, fRoles[i])
	}
	for i := 0; i < len(fCalls); i++ {
		t := strings.Split(fCalls[i], ",")
		fCallsFinal = append(fCallsFinal, t)
	}

	return fCallsFinal, fRolesFinal, sID
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
