package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, _ := readFile("inputs.txt")
	trebuchet(file)
}

func trebuchet(input []string) {
	total := 0
	for i := 0; i < len(input); i++ {
		transformedInput := processString(input[i])
		firstNumber := ""
		lastNumber := ""
		for _, ch := range transformedInput {
			if unicode.IsDigit(ch) {
				if firstNumber == "" {
					firstNumber = string(ch)
				} else {
					lastNumber = string(ch)
				}
			}
		}
		if lastNumber == "" {
			lastNumber = firstNumber
		}
		number, _ := strconv.Atoi(firstNumber + lastNumber)
		total = total + number
	}
	fmt.Println(total)
}

func processString(input string) string {
	numbersMap := map[string]string{
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
	var processedString string
	var currentToken string
	var nextToken string
	for i, char := range input {
		if unicode.IsLetter(char) {
			currentToken += string(char)
			if nextToken != "" {
				nextToken += string(char)
			}
			for i := range numbersMap {
				if strings.Contains(currentToken, i) {
					processedString += numbersMap[i]
					nextToken = currentToken[len(currentToken)-1:]
					currentToken = ""
				}
				if strings.Contains(nextToken, i) {
					processedString += numbersMap[i]
					nextToken = ""
					currentToken = ""
				}
			}

		} else {
			processedString += currentToken
			processedString += string(char)
			currentToken = ""
			nextToken = ""
		}
		if i == len(input)-1 {
			processedString += currentToken
			processedString += string(char)
			currentToken = ""
			nextToken = ""
		}
	}
	return processedString
}

func readFile(file string) ([]string, error) {
	var inputs []string

	f, e := os.Open(file)
	if e != nil {
		return nil, e
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	if e := scanner.Err(); e != nil {
		return nil, e
	}

	return inputs, nil
}
