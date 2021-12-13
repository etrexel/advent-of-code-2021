package solutions

import "testing"

var input = []int{3, 4, 3, 1, 2}

func TestDaySixPartOne(t *testing.T) {
	expected := 5934
	result := DaySixPartOne(input)
	if result != expected {
		t.Fatalf("Expected: %d, got: %d", expected, result)
	}
}

func TestDayTwoSixTwo(t *testing.T) {
	expected := 26984457539
	result := DaySixPartTwo(input)
	if result != expected {
		t.Fatalf("Expected: %d, got: %d", expected, result)
	}
}

func TestTotalFish(t *testing.T) {
	cache := make(map[int]int)
	expected := 5
	result := totalFish(3, 18, cache)
	if result != expected {
		t.Fatalf("Count: %d, expected: %d, got: %d\n", 3, expected, result)
	}
	expected = 4
	result = totalFish(4, 18, cache)
	if result != expected {
		t.Fatalf("Count: %d, expected: %d, got: %d\n", 4, expected, result)
	}
	expected = 7
	result = totalFish(1, 18, cache)
	if result != expected {
		t.Fatalf("Count: %d, expected: %d, got: %d\n", 1, expected, result)
	}
	expected = 5
	result = totalFish(2, 18, cache)
	if result != expected {
		t.Fatalf("Count: %d, expected: %d, got: %d\n", 2, expected, result)
	}
	expected = 26
	result = 0
	for _, val := range input {
		result += totalFish(val, 18, cache)
	}
	if result != expected {
		t.Fatalf("Expected: %d, got: %d\n", expected, result)
	}
}
