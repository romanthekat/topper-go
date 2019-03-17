package main

import (
	"testing"
)

func TestParsingZshLine(t *testing.T) {
	rawLine := ": 000000000:0;vim ~/.zshrc"

	command := getCommandFromZshHistoryLine(rawLine)

	if command != "vim ~/.zshrc" {
		t.Errorf("Command got wrong from zsh line format.\n"+
			"rawLine:%s\n"+
			"command:%s\n",
			rawLine, command)
	}
}
