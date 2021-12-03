package solutions

import (
	"math"
	"os"
	"strings"
)

func DayThreeReadInput(filePath string) ([]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(data), "\n"), nil
}

func DayThreePartOne(input []string) int {
	gamma := getGamma(input)
	epsilon := getEpsilon(input)
	return gamma * epsilon
}

func DayThreePartTwo(input []string) int {
	oxygen := getOxygen(input)
	co2 := getCO2(input)
	return oxygen * co2
}

func bitsToInt(input string) int {
	result := 0
	power := 0
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] == '1' {
			result += int(math.Pow(float64(2), float64(power)))
		}
		power++
	}
	return result
}

func mostCommonBit(input []string, idx int) int {
	zeroCount, oneCount := 0, 0
	for _, val := range input {
		switch val[idx] {
		case '0':
			zeroCount++
			break
		case '1':
			oneCount++
			break
		}
	}
	if zeroCount > oneCount {
		return 0
	}
	return 1
}

func getGamma(input []string) int {
	gamma := 0
	bit := len(input[0]) - 1
	power := 0
	for bit >= 0 {
		currBit := mostCommonBit(input, bit)
		if currBit == 1 {
			gamma += int(math.Pow(float64(2), float64(power)))
		}
		bit--
		power++
	}
	return gamma
}

func getEpsilon(input []string) int {
	epsilon := 0
	bit := len(input[0]) - 1
	power := 0
	for bit >= 0 {
		currBit := mostCommonBit(input, bit)
		if currBit == 0 {
			epsilon += int(math.Pow(float64(2), float64(power)))
		}
		bit--
		power++
	}
	return epsilon
}

func bitFilter(input []string, idx int, filterType string) []string {
	resultLen := 0
	bit := mostCommonBit(input, idx)
	filterVal := '0'
	if filterType == "MOST" {
		if bit == 0 {
			filterVal = '0'
		} else {
			filterVal = '1'
		}
	} else {
		if bit == 1 {
			filterVal = '0'
		} else {
			filterVal = '1'
		}
	}
	for _, val := range input {
		check := rune(val[idx])
		if check == filterVal {
			resultLen++
		}
	}
	result := make([]string, resultLen)
	i := 0
	for _, val := range input {
		check := rune(val[idx])
		if check == filterVal {
			result[i] = val
			i++
		}
	}
	return result
}

func getOxygen(input []string) int {
	filterType := "MOST"
	result := bitFilter(input, 0, filterType)
	idx := 1
	for len(result) > 1 {
		result = bitFilter(result, idx, filterType)
		idx++
	}
	return bitsToInt(result[0])
}

func getCO2(input []string) int {
	filterType := "LEAST"
	result := bitFilter(input, 0, filterType)
	idx := 1
	for len(result) > 1 {
		result = bitFilter(result, idx, filterType)
		idx++
	}
	return bitsToInt(result[0])
}
