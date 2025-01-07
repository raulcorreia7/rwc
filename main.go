package main

import (
	"fmt"
	"os"
)

/*
This is a crude implementation of `wc` to learn the golanguage.

source: https://codingchallenges.fyi/challenges/challenge-wc
manpage: https://linux.die.net/man/1/wc
@author: r.correia
*/

type FileStatistics struct {
	Bytes int
	Lines int
	Words int
}

func main() {
	args := os.Args[1:]

	if len(args) > 2 {
		fmt.Println("Too many arguments")
	}

	command, filename := ProcessArguments(args)

	var file *os.File = OpenFile(filename)
	defer file.Close()

	file_stats := ProcessCommand(command, file)

	WriteOutput(command, file_stats, filename)
}

func WriteOutput(command string, file_stats FileStatistics, filename string) {
	output := ""
	if command == "" {
		output = fmt.Sprintf("  %d  %d %d %s", file_stats.Lines, file_stats.Words, file_stats.Bytes, filename)
	}

	if command == "-c" {
		output = fmt.Sprintf("%d %s", file_stats.Bytes, filename)
	} else if command == "-l" {
		output = fmt.Sprintf("%d %s", file_stats.Lines, filename)
	} else if command == "-w" {
		output = fmt.Sprintf("%d %s", file_stats.Words, filename)
	}

	fmt.Println(output)
}

func OpenFile(filename string) *os.File {
	var file *os.File

	// if no file specified, use Stdin
	if len(filename) == 0 {
		file = os.Stdin
	} else {
		temp, err := os.Open(filename)
		if err != nil {
			err = fmt.Errorf("error opening file: %v", err)
			fmt.Println(err)
		}
		file = temp
	}

	return file
}
