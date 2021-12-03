package solutions

import (
	"os"
	"strconv"
	"strings"
)

func DayOneReadInput(filePath string) ([]int, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	values := strings.Split(string(data), "\n")
	input := make([]int, len(values))
	for idx, value := range values {
		i, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		input[idx] = i
	}
	return input, nil
}

func DayOnePartOne(input []int) int {
	previousDepth := input[0]
	count := 0
	for _, depth := range input[1:] {
		if depth > previousDepth {
			count++
		}
		previousDepth = depth
	}
	return count
}

func DayOnePartTwo(input []int) int {
	prevWindow := input[0] + input[1] + input[2]
	count := 0
	for i := 1; i < len(input) - 1; i++ {
		currWindow := input[i - 1] + input[i] + input[i + 1]
		if currWindow > prevWindow {
			count++
		}
		prevWindow = currWindow
	}
	return count
}
