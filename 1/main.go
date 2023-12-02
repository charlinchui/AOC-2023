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

/*
--- Day 1: Trebuchet?! ---
Something is wrong with global snow production, and you've been selected to take a look. The Elves have even given you a map; on it, they've used stars to mark the top fifty locations that are likely to be having problems.

You've been doing this long enough to know that to restore snow operations, you need to check all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

You try to ask why they can't just use a weather machine ("not powerful enough") and where they're even sending you ("the sky") and why your map looks mostly blank ("you sure ask a lot of questions") and hang on did you just say the sky ("of course, where do you think snow comes from") when you realize that the Elves are already loading you into a trebuchet ("please hold still, we need to strap you in").

As they're making the final adjustments, they discover that their calibration document (your puzzle input) has been amended by a very young Elf who was apparently just excited to show off her art skills. Consequently, the Elves are having trouble reading the values on the document.

The newly-improved calibration document consists of lines of text; each line originally contained a specific calibration value that the Elves now need to recover. On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.

For example:

1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet

In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.

Consider your entire calibration document. What is the sum of all of the calibration values?

--- Part Two ---
Your calculation isn't quite right. It looks like some of the digits are actually spelled out with letters: one, two, three, four, five, six, seven, eight, and nine also count as valid "digits".

Equipped with this new information, you now need to find the real first and last digit on each line. For example:

two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
In this example, the calibration values are 29, 83, 13, 24, 42, 14, and 76. Adding these together produces 281.

What is the sum of all of the calibration values?
*/
