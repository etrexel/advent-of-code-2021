package solutions

import "testing"

func TestDayTwoPartOne(t *testing.T) {
	input := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	expected := 150
	result, err := DayTwoPartOne(input)
	if err != nil {
		t.Fatalf("Error in function: %s", err)
	}
	if result != expected {
		t.Fatalf("Expected: %d, got: %d", expected, result)
	}
}

func TestDayTwoPartTwo(t *testing.T) {
	input := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	expected := 900
	result, err := DayTwoPartTwo(input)
	if err != nil {
		t.Fatalf("Error in function: %s", err)
	}
	if result != expected {
		t.Fatalf("Expected: %d, got: %d", expected, result)
	}
}
