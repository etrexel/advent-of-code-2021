package solutions

import "testing"

var daySevenTestInput = []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

func TestDaySevenPartOne(t *testing.T) {
	expected := 37
	result := DaySevenPartOne(daySevenTestInput)
	if result != expected {
		t.Fatalf("Expected: %d, got: %d\n", expected, result)
	}
}

func TestDaySevenPartTwo(t *testing.T) {
	expected := 168
	result := DaySevenPartTwo(daySevenTestInput)
	if result != expected {
		t.Fatalf("Expected: %d, got: %d\n", expected, result)
	}
}

func TestCalculateCost(t *testing.T) {
	expected := 6
	result := calculateCost(3)
	if result != expected {
		t.Fatalf("Expected: %d, got: %d\n", expected, result)
	}
	expected = 10
	result = calculateCost(4)
	if result != expected {
		t.Fatalf("Expected: %d, got: %d\n", expected, result)
	}
}
