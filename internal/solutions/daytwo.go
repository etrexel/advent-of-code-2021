package solutions

import (
	"os"
	"strconv"
	"strings"
)

func DayTwoReadInput(filePath string) ([]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(data), "\n"), nil
}

func DayTwoPartOne(input []string) (int, error) {
	horizontal, vertical := 0, 0
	for _, line := range input {
		components := strings.Split(line, " ")
		delta, err := strconv.Atoi(components[1])
		if err != nil {
			return 0, err
		}
		switch components[0] {
		case "forward":
			horizontal += delta
			break
		case "down":
			vertical += delta
			break
		case "up":
			vertical -= delta
			break
		}
	}
	return horizontal * vertical, nil
}

func DayTwoPartTwo(input []string) (int, error) {
	horizontal, vertical, aim := 0, 0, 0
	for _, line := range input {
		components := strings.Split(line, " ")
		delta, err := strconv.Atoi(components[1])
		if err != nil {
			return 0, err
		}
		switch components[0] {
		case "forward":
			horizontal += delta
			vertical += delta * aim
			break
		case "down":
			aim += delta
			break
		case "up":
			aim -= delta
			break
		}
	}
	return horizontal * vertical, nil
}
