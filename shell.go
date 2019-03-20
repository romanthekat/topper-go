package main

import (
	"fmt"
	"log"
	"os/user"
	"strings"
)

//Shell contains information enough to define shell and parse its history file
type Shell struct {
	binaryName           string
	historyFilename      string
	historyLineToCommand func(rawLine string) string
}

func (s Shell) String() string {
	return fmt.Sprintf("%s %s", s.binaryName, s.historyFilename)
}

func (s Shell) getHistoryFullFilename() string {
	homeDir := getUserHomeDir()

	return fmt.Sprintf("%s/%s", homeDir, s.historyFilename)
}

//SupportedShells contains descriptions how to handle different shells types
var SupportedShells = []Shell{
	{binaryName: "bash", historyFilename: ".bash_history", historyLineToCommand: returnAsIs},
	{binaryName: "zsh", historyFilename: ".zsh_history", historyLineToCommand: getCommandFromZshHistoryLine},
}

var unknownShell = Shell{"unknown shell", "unknown history file", nil}

func getShellByBinary(shellBinary string) Shell {
	for _, shell := range SupportedShells {
		if strings.Contains(shellBinary, shell.binaryName) {
			return shell
		}
	}

	return unknownShell
}

var returnAsIs = func(rawLine string) string {
	return rawLine
}

//extended history file format is in use
//‘: <beginning time>:<elapsed seconds>;<command>’
var getCommandFromZshHistoryLine = func(rawLine string) string {
	delimiterPos := strings.Index(rawLine, ";")
	if delimiterPos == -1 {
		return "" //skip
	}

	return rawLine[delimiterPos+1:]
}

func getUserHomeDir() string {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return currentUser.HomeDir
}

func equals(first, second Shell) bool {
	return first.binaryName == second.binaryName
}
