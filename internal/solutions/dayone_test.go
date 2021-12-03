package solutions

import "testing"

func TestDayOnePartOne(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	expected := 7
	result := DayOnePartOne(input)
	if result != expected {
		t.Fatalf("Expected: %d, got: %d", expected, result)
	}
}

func TestDayOnePartTwo(t *testing.T) {
	input := []int{607, 618, 618, 617, 647, 716, 769, 792}
	expected := 5
	result := DayOnePartTwo(input)
	if result != expected {
		t.Fatalf("Expected: %d, got: %d", expected, result)
	}
}
