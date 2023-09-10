package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	content, err := os.Open("books/Frankenstein.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer content.Close()
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
