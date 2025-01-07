package main

import "os"

func ProcessCommand(command string, file *os.File) FileStatistics {
	var file_stats FileStatistics
	switch command {
	case "-c":
		{
			file_stats = Count(file, Bytes)
		}
	case "-l":
		{
			file_stats = Count(file, Lines)
		}
	case "-w":
		{
			file_stats = Count(file, Words)
		}
	default:
		{
			file_stats = Count(file, Bytes|Lines|Words)
		}
	}
	return file_stats
}
