package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var (
	wordToNumber = map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
)

func replaceNumberWords(s string) string {
	var current string
	idxCount := 0
	for i := 0; i < len(s); i++ {
		for k, v := range wordToNumber {
			if strings.HasPrefix(s[i:], k) {
				current += v
				idxCount = len(k)
				break
			}
		}

		if idxCount > 0 {
			idxCount--
			continue
		}
		current += string(s[i])
	}
	return current
}

func main() {
	readFile, err := os.Open("./one.txt")

	if err != nil {
		log.Fatal("Unable to open input file")
	}

	defer readFile.Close()

	fScanner := bufio.NewScanner(readFile)

	sum := 0

	for fScanner.Scan() {
		line := fScanner.Text()
		line = replaceNumberWords(line)
		var first string
		var last string
		lastIndex := len(line) - 1
		i := 0
		for first == "" || last == "" {
			if first == "" && unicode.IsDigit(rune(line[i])) {
				first = string(line[i])
			}

			if last == "" && unicode.IsDigit(rune(line[lastIndex-i])) {
				last = string(line[lastIndex-i])
			}
			i++
		}
		if v, cErr := strconv.Atoi(first + last); cErr == nil {
			sum += v
		}
	}

	log.Println(sum)
}
