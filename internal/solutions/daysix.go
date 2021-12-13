package solutions

import (
	"os"
	"strconv"
	"strings"
)

func DaySixReadInput(filePath string) ([]int, error) {
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

func DaySixPartOne(input []int) int {
	return simulateDays(input, 80)
}

func DaySixPartTwo(input []int) int {
	return simulateDays(input, 256)
}

func simulateDays(input []int, days int) int {
	cache := make(map[int]int)
	total := 0
	for _, val := range input {
		total += totalFish(val, days, cache)
	}
	return total
}

func totalFish(count int, days int, cache map[int]int) int {
	if count == 8 && cache[days] != 0 {
		return cache[days]
	}
	total := 1
	daysLeft := days
	if days < count+1 {
		return total
	}
	daysLeft -= count + 1
	total += totalFish(8, daysLeft, cache)
	for i := 1; i <= daysLeft/7; i++ {
		total += totalFish(8, daysLeft-(7*i), cache)
	}
	if count == 8 && cache[days] == 0 {
		cache[days] = total
	}
	return total
}
