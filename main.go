package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"os/user"
)

func main() {
	commandsFreq := getCommandsFreq()

	fmt.Println(commandsFreq)
}

func getCommandsFreq() map[string]int {
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