package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func GameScore(game string, maxScores map[string]int) (int, int) {
	gameName, cubeCounts, _ := strings.Cut(game, ":")
	_, gID, _ := strings.Cut(gameName, " ")

	re := regexp.MustCompile(`(\d+) (red|green|blue)`)

	matches := re.FindAllStringSubmatch(cubeCounts, -1)

	valid := true
	score := 0
	// r, g, b
	mins := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, m := range matches {
		minNeeded, ok := mins[m[2]]

		if !ok {
			continue
		}

		if comp, err := strconv.Atoi(m[1]); err == nil {
			if comp > minNeeded {
				mins[m[2]] = comp
			}

			if valid {
				maxAllowed, sok := maxScores[m[2]]
				if sok && comp > maxAllowed {
					valid = false
				}
			}
		}

	}

	if valid {
		if gIDValue, err := strconv.Atoi(gID); err == nil {
			score = gIDValue
		}
	}

	return score, mins["red"] * mins["green"] * mins["blue"]
}

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal("Unable to open input file")
	}

	defer readFile.Close()

	fScanner := bufio.NewScanner(readFile)

	sum := 0
	powerSum := 0

	maxCounts := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for fScanner.Scan() {
		line := fScanner.Text()
		score, powers := GameScore(line, maxCounts)
		sum += score
		powerSum += powers
	}

	log.Println(sum, powerSum)
}
