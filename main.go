package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.Open("books/Frankenstein.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer content.Close()
	fmt.Println(charCount(content))
	wordCount := wordCount(content)

	fmt.Println("The number of words is :", wordCount)
}

func wordCount(text io.Reader) int {
	count := 0

	scanner := bufio.NewScanner(text)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		count++
	}

	return count
}

func charCount(text io.Reader) map[string]int {
	counts := map[string]int{}

	scanner := bufio.NewScanner(text)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		char := strings.ToLower(scanner.Text())
		counts[char]++
	}
	return counts
}
