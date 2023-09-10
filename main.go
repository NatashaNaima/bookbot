package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
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

	wordCount := wordCount(data)
	letterCount := charCount(data)
	keys := make([]string, 0, len(letterCount))
	for k := range letterCount {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return letterCount[keys[i]] > letterCount[keys[j]]
	})

	fmt.Println("The number of words is :", wordCount)

	for _, k := range keys {
		fmt.Println(k, letterCount[k])
	}
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
