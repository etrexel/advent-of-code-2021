package solutions

import "testing"

var inputArr = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
}

func TestBitsToInt(t *testing.T) {
	input, expected := "1010", 10
	result := bitsToInt(input)
	if result != expected {
		t.Fatalf("Expected: %d, got: %d\n", expected, result)
	}
	input, expected = "111", 7
	result = bitsToInt(input)
	if result != expected {
		t.Fatalf("Expected: %d, got: %d\n", expected, result)
	}
}

func TestMostCommonBit(t *testing.T) {
	idx, expected := 0, 1
	result := mostCommonBit(inputArr, idx)
	if result != expected {
		t.Fatalf("Index: %d, expected %d, got: %d\n", idx, expected, result)
	}
	idx, expected = 1, 0
	result = mostCommonBit(inputArr, idx)
	if result != expected {
		t.Fatalf("Index: %d, expected %d, got: %d\n", idx, expected, result)
	}
}

func TestGetGamma(t *testing.T) {
	expected := 22
	result := getGamma(inputArr)
	if result != expected {
		t.Fatalf("Expected: %d, got: %d\n", expected, result)
	}
}

func TestGetEpsilon(t *testing.T) {
	expected := 9
	result := getEpsilon(inputArr)
	if result != expected {
		t.Fatalf("Expected: %d, got: %d\n", expected, result)
	}
}

func TestBitFilter(t *testing.T) {
	expected := []string {
		"11110",
		"10110",
		"10111",
		"10101",
		"11100",
		"10000",
		"11001",
	}
	result := bitFilter(inputArr, 0, "MOST")
	if len(expected) != len(result) {
		t.Fatalf("Lengths differ, expected: %d, got: %d\n", len(expected), len(result))
	}
	for i, v := range expected {
		if v != result[i] {
			t.Fatalf("Value differs, expected: %s, got %s\n", v, result[i])
		}
	}
}

func TestGetOxygen(t *testing.T) {
	expected := 23
	result := getOxygen(inputArr)
	if result != expected {
		t.Fatalf("Expected: %d, got: %d\n", expected, result)
	}
}

func TestGetCO2(t *testing.T) {
	expected := 10
	result := getCO2(inputArr)
	if result != expected {
		t.Fatalf("Expected: %d, got: %d\n", expected, result)
	}
}

func TestDayThreePartOne(t *testing.T) {
	expected := 198
	result := DayThreePartOne(inputArr)
	if result != expected {
		t.Fatalf("Expected: %d, got: %d\n", expected, result)
	}
}

func TestDayThreePartTwo(t *testing.T) {
	expected := 230
	result := DayThreePartTwo(inputArr)
	if result != expected {
		t.Fatalf("Expected: %d, got: %d\n", expected, result)
	}
}
