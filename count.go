package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CountType uint8

const (
	Bytes CountType = 1 << iota
	Lines
	Words
)

func Count(file *os.File, flag CountType) FileStatistics {
	file_stats := FileStatistics{
		Bytes: 0,
		Lines: 0,
		Words: 0,
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if (flag & Bytes) > 0 {
			file_stats.Bytes += int(len(scanner.Bytes()))
		}
		if (flag & Lines) > 0 {
			file_stats.Lines++
		}
		if (flag & Words) > 0 {
			file_stats.Words += len(strings.Fields(scanner.Text()))
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error:", err)
	}

	return file_stats
}
