package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"os/user"
	"sort"
)

type Command struct {
	command string
	number int
	freq int
}

type Commands []*Command

func (slice Commands) Len() int {
	return len(slice)
}

func (slice Commands) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (slice Commands) Less(i, j int) bool {
	return slice[i].freq < slice[j].freq
}


func main() {
	commandsList := getCommands()
	sort.Sort(sort.Reverse(Commands(commandsList)))

	for _, command := range commandsList {
		fmt.Printf("%5d: %v (x%d)\n", command.number, command.command, command.freq)
	}
}

func getCommands() []*Command {
	return getCommandsByValues(getHistoryContent())
}

func getCommandsByValues(commandsChan <-chan string) []*Command {
	commandStructs := make(map[string]*Command)

	number := 1
	for commandString := range commandsChan {
		command, exists := commandStructs[commandString]
		if exists {
			command.freq = command.freq + 1
		} else {
			commandStructs[commandString] = &Command{command: commandString, number: number, freq: 1}
		}

		number++
	}

	return getValuesFromMap(commandStructs)
}

func getValuesFromMap(commandStructs map[string]*Command) []*Command {
	values := make([]*Command, 0, len(commandStructs))

	for _, value := range commandStructs {
		values = append(values, value)
	}

	return values
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