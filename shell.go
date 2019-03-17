package main

import (
	"log"
	"strings"
)

type Shell struct {
	binaryName     string
	historyHandler func(rawLine string) string
}

var shells = []Shell{
	Shell{binaryName: "bash", historyHandler: returnAsIs},
	Shell{"zsh", nil},
}

var returnAsIs = func(rawLine string) string {
	return rawLine
}

//extended history file format is in use
//‘: <beginning time>:<elapsed seconds>;<command>’
var getCommandFromZshHistoryLine = func (rawLine string) string {
	delimiterPos := strings.Index(rawLine, ";")
	if delimiterPos == -1 {
		log.Fatalf("zsh history line parsing failed, searching for ';' failed, line: %s", rawLine)
	}

	return rawLine[delimiterPos+ 1:]
}