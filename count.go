package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CountBytes(file *os.File) int {
	file.Seek(0, 0)
	bytes, err := file.Stat()

	if err != nil {
		err = fmt.Errorf("error counting bytes: %v", err)
		fmt.Println(err)
	}

	return int(bytes.Size())

}

func CountLines(file *os.File) int {
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error:", err)
	}
	return int(count)
}

func CountWords(file *os.File) int {
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	words := 0

	for scanner.Scan() {
		text := scanner.Text()
		words += len(strings.Fields(text))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error:", err)
	}
	return int(words)
}

func CountCharacters(file *os.File) int {
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	characters := 0

	for scanner.Scan() {
		characters += len(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error:", err)
	}

	return int(characters)
}
