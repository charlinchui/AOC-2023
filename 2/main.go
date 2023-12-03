package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputs, _ := readFile("inputs.txt")
	//fmt.Println(inputs)
	fmt.Println(cubeCanondrom(inputs))
}

func cubeCanondrom(inputs []string) int {
	power := 0
	possibleGames := 0
	for i := 0; i < len(inputs); i++ {
		games := strings.Split(strings.Split(inputs[i], ":")[1], ";")
		minBlue := 0
		minGreen := 0
		minRed := 0
		for i := 0; i < len(games); i++ {
			rounds := strings.Split(games[i], ",")
			for i := range rounds {
				value, _ := strconv.Atoi(strings.Split(rounds[i], " ")[1])
				color := strings.Split(rounds[i], " ")[2]
				if color == "green" && value > minGreen {
					minGreen = value
				}
				if color == "blue" && value > minBlue {
					minBlue = value
				}
				if color == "red" && value > minRed {
					minRed = value
				}
			}
			fmt.Println("b", minBlue, "g", minGreen, "r", minRed)
			power = minBlue * minGreen * minRed
		}
		possibleGames += power
	}

	return possibleGames
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
