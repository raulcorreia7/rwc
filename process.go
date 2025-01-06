package main

import "os"

func ProcessCommand(command string, file *os.File) FileStatistics {
	file_stats := FileStatistics{
		Bytes: -1,
		Lines: -1,
		Words: -1,
	}

	switch command {
	case "-c":
		{
			file_stats.Bytes = CountBytes(file)
		}
	case "-l":
		{
			file_stats.Lines = CountLines(file)
		}
	case "-w":
		{
			file_stats.Words = CountWords(file)
		}
	default:
		{
			file_stats.Bytes = CountBytes(file)
			file_stats.Lines = CountLines(file)
			file_stats.Words = CountWords(file)
		}
	}

	return file_stats
}
