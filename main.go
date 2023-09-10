package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	content, err := os.Open("books/Frankenstein.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer content.Close()
	// Read the content of the file into memory
	data := make([]byte, 0)
	buffer := make([]byte, 1024) // You can adjust the buffer size as needed

	for {
		n, err := content.Read(buffer)
		if err != nil && n == 0 {
			break
		}
		data = append(data, buffer[:n]...)
	}

	fmt.Println(charCount(data))

	wordCount := wordCount(data)
	fmt.Println("The number of words is :", wordCount)
}

func wordCount(data []byte) int {
	count := 0

	text := bytes.NewReader(data)

	scanner := bufio.NewScanner(text)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		count++
	}

	return count
}

func charCount(data []byte) map[string]int {
	counts := map[string]int{}

	text := bytes.NewReader(data)

	scanner := bufio.NewScanner(text)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		char := strings.ToLower(scanner.Text())
		if unicode.IsLetter(rune(char[0])) {
			counts[char]++
		}
	}
	return counts
}
