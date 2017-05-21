package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"os/user"
	"sort"
)

type command struct {
	command string
	freq int
}

type commands []command

func (slice commands) Len() int {
	return len(slice)
}
func (slice commands) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
func (slice commands) Less(i, j int) bool {
	return slice[i].freq < slice[j].freq
}

func main() {
	commandsList := initCommands(getCommandsFrequencies())
	sort.Sort(commands(commandsList))

	fmt.Println(commandsList)
}

func initCommands(commandsFreq map[string]int) []command {
	commands := []command{}

	for cmd, freq := range commandsFreq {
		commands = append(commands, command{cmd, freq})
	}

	return commands
}

func getCommandsFrequencies() map[string]int {
	commandsFreq := make(map[string]int)

	for command := range getHistoryContent() {
		value, exists := commandsFreq[command]
		if exists {
			commandsFreq[command] = value + 1
		} else {
			commandsFreq[command] = 1
		}
	}

	return commandsFreq
}

func getHistoryContent() <-chan string {
	historyFilename := getHistoryFilename()
	return ReadByLine(historyFilename)
}

func getHistoryFilename() string {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return currentUser.HomeDir + "/.bash_history"
}


func ReadByLine(filename string) <-chan string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := make(chan string)

	go func() {
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for (scanner.Scan()) {
			lines <- scanner.Text()
		}

		close(lines)

		scannerErr := scanner.Err()
		if scannerErr != nil {
			log.Fatal(scannerErr)
		}
	}()

	return lines
}