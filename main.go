package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

const defaultTopCommandsCount = 10

//Command contains shell command information within all calls
type Command struct {
	command string
	number  int
	freq    int
}

func (c Command) String() string {
	return fmt.Sprintf("%5d: %v (x%d)", c.number, c.command, c.freq)
}

//Commands is sortable (by freq) collection of shell commands
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

//CommandsCalls is a channel of raw commands sequence
type CommandsCalls <-chan string

func main() {
	topCommandsCount := getTopCommandsCount(defaultTopCommandsCount)
	shellHistory := getShellHistory()

	topCommands := getTopCommands(shellHistory, topCommandsCount)

	for _, command := range topCommands {
		fmt.Println(command)
	}
}

func getShellHistory() CommandsCalls {
	shell := getCurrentShell()
	return readShellHistory(shell)
}

func getTopCommands(shellHistory CommandsCalls, topCommandsCount int) Commands {
	commands := getCommands(shellHistory)
	sort.Sort(sort.Reverse(commands))

	return commands[0:min(topCommandsCount, len(commands))]
}

//TODO command number calculation is a bit obscure
func getCommands(shellHistory CommandsCalls) Commands {
	commands := make(map[string]*Command)

	number := 1
	for commandString := range shellHistory {
		command, exists := commands[commandString]
		if exists {
			command.freq = command.freq + 1
			command.number = number
		} else {
			commands[commandString] = &Command{command: commandString, number: number, freq: 1}
		}

		number++
	}

	return getValuesFromMap(commands)
}

func getCurrentShell() Shell {
	shellBinary := os.Getenv("SHELL")

	return getCurrentShellByBinary(shellBinary)
}

func getCurrentShellByBinary(shellBinary string) Shell {
	shell := getShellByBinary(shellBinary)

	if equals(shell, unknownShell) {
		log.Fatalf("Unknown shell detected: %s", shellBinary)
	}

	return shell
}

func getTopCommandsCount(defaultTopCommandsCount int) int {
	args := os.Args
	if len(args) > 1 {
		topCommandsCount, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal(err)
		}

		return topCommandsCount
	} else {
		return defaultTopCommandsCount
	}
}

func getValuesFromMap(commandStructs map[string]*Command) Commands {
	values := make([]*Command, 0, len(commandStructs))

	for _, value := range commandStructs {
		values = append(values, value)
	}

	return values
}

func readShellHistory(shell Shell) CommandsCalls {
	file, err := os.Open(shell.getHistoryFullFilename())
	if err != nil {
		log.Fatal(err)
	}

	lines := make(chan string)

	go func() {
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			line = shell.historyLineToCommand(line)

			if lineContainsCommand(line) {
				lines <- line
			}
		}

		close(lines)

		scannerErr := scanner.Err()
		if scannerErr != nil {
			log.Fatal(scannerErr)
		}
	}()

	return lines
}

func lineContainsCommand(line string) bool {
	return len(line) > 0
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
