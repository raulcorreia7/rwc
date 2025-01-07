package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
)

func ProcessArguments(args []string) (string, string) {
	command := ""
	filename := ""

	supported_commands := []string{"-c", "-l", "-w"}
	for _, arg := range args {

		if arg[0] == '-' {
			if !slices.Contains(supported_commands, arg) {
				fmt.Printf("%s: unknown options -- %s", filepath.Base(os.Args[0]), arg[1:])
				os.Exit(1)
			} else {
				command = arg
			}
		} else {
			filename = arg
		}
	}
	return command, filename
}
