package solutions

import (
	"math"
	"os"
	"strconv"
	"strings"
)

func DaySevenReadInput(filePath string) ([]int, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	values := strings.Split(string(data), ",")
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

func DaySevenPartOne(input []int) int {
	costFunc := func(cost int) int { return cost }
	return lowestCost(input, costFunc)
}

func DaySevenPartTwo(input []int) int {
	return lowestCost(input, calculateCost)
}

func lowestCost(input []int, costFunction func(int) int) int {
	max := 0
	for _, val := range input {
		if val >= max {
			max = val
		}
	}
	cost := math.MaxInt
	for i := 0; i < max; i++ {
		currentCost := 0
		for _, val := range input {
			currentCost += costFunction(int(math.Abs(float64(i - val))))
		}
		if currentCost <= cost {
			cost = currentCost
		}
	}
	return cost
}

func calculateCost(distance int) int {
	result := (distance / 2) * (distance + 1)
	if distance%2 != 0 {
		result += (distance + 1) / 2
	}
	return result
}
