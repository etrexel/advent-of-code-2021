package solutions

import (
	"testing"

	"github.com/etrexel/advent-of-code-2021/pkg/types"
)

var pointPairs = []types.PointPair{
	{
		types.Point{X: 0, Y: 9},
		types.Point{X: 5, Y: 9},
	},
	{
		types.Point{X: 8, Y: 0},
		types.Point{X: 0, Y: 8},
	},
	{
		types.Point{X: 9, Y: 4},
		types.Point{X: 3, Y: 4},
	},
	{
		types.Point{X: 2, Y: 2},
		types.Point{X: 2, Y: 1},
	},
	{
		types.Point{X: 7, Y: 0},
		types.Point{X: 7, Y: 4},
	},
	{
		types.Point{X: 6, Y: 4},
		types.Point{X: 2, Y: 0},
	},
	{
		types.Point{X: 0, Y: 9},
		types.Point{X: 2, Y: 9},
	},
	{
		types.Point{X: 3, Y: 4},
		types.Point{X: 1, Y: 4},
	},
	{
		types.Point{X: 0, Y: 0},
		types.Point{X: 8, Y: 8},
	},
	{
		types.Point{X: 5, Y: 5},
		types.Point{X: 8, Y: 2},
	},
}

func TestDayFivePartOne(t *testing.T) {
	expected := 5
	result := DayFivePartOne(pointPairs)
	if result != expected {
		t.Fatalf("Expected: %d, got: %d", expected, result)
	}
}

func TestDayFivePartTwo(t *testing.T) {
	expected := 12
	result := DayFivePartTwo(pointPairs)
	if result != expected {
		t.Fatalf("Expected: %d, got: %d", expected, result)
	}
}
