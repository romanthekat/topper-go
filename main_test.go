package main

import (
	"strings"
	"testing"
)

func TestGetTopCommandsGenericUsecase(t *testing.T) {
	shellCommands := getMockShellCommands()
	topCommandsCount := 3

	commands := getTopCommands(shellCommands, topCommandsCount)

	if len(commands) != topCommandsCount {
		t.Errorf("Commands count differs from %d, commands = %s", topCommandsCount, commands)
	}

	commandThree := commands[0]
	if !strings.Contains(commandThree.command, "three") {
		t.Errorf("Wrong command, should be three, but %s", commandThree)
	}

	commandTwo := commands[1]
	if !strings.Contains(commandTwo.command, "two") {
		t.Errorf("Wrong command, should be two, but %s", commandTwo)
	}

	commandOne := commands[2]
	if !strings.Contains(commandOne.command, "one") {
		t.Errorf("Wrong command, should be one, but %s", commandOne)
	}
}

func TestTopCommandsCountIfFoundLess(t *testing.T) {
	shellCommands := getMockShellCommands() //returns 6 commands, 3 different
	topCommandsCount := 4                   //more than actually provided

	commands := getTopCommands(shellCommands, topCommandsCount)

	if len(commands) != 3 {
		t.Errorf("Only 3 different commands provided, "+
			"requested top count is %d, "+
			"commands size must be 3\n"+
			"commands = %s",
			topCommandsCount, commands)
	}
}

func TestTopCommandsCountIfFoundMore(t *testing.T) {
	shellCommands := getMockShellCommands() //returns 6 commands, 3 different
	topCommandsCount := 2                   //less than actually provided

	commands := getTopCommands(shellCommands, topCommandsCount)

	if len(commands) != topCommandsCount {
		t.Errorf("Only 3 different commands provided, "+
			"requested top count is %d, which is less,"+
			"commands size must be %d\n"+
			"commands = %s",
			topCommandsCount, topCommandsCount, commands)
	}
}

func TestGetShellByBinary(t *testing.T) {
	shell := getShellByBinary("/bin/bash")

	if shell.binaryName != "bash" {
		t.Fatalf("Wrong shell found, must be bash, but %s", shell)
	}
}

func TestGetShellByBinaryUnknownShell(t *testing.T) {
	shell := getShellByBinary("/meow/shell")

	if shell.binaryName != unknownShell.binaryName {
		t.Fatalf("Wrong shell found, must be unknown, but %s", shell)
	}
}

func getMockShellCommands() <-chan string {
	shellCommands := make(chan string, 6)

	shellCommands <- "one"
	shellCommands <- "two"
	shellCommands <- "two"
	shellCommands <- "three"
	shellCommands <- "three"
	shellCommands <- "three"

	close(shellCommands)

	return shellCommands
}
