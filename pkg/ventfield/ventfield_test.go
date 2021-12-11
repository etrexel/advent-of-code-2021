package ventfield

import (
	"testing"

	"github.com/etrexel/advent-of-code-2021/pkg/types"
)

func TestAddLine(t *testing.T) {
	vf := NewVentField(5, 5)
	pair := types.PointPair{{2, 1}, {2, 4}}
	vf.AddLine(pair)
	for i := 1; i < 4; i++ {
		if vf.field[i][2] != 1 {
			t.Fatalf("Failed to add line\n")
		}
	}
}

func TestCrossingPointCount(t *testing.T) {
	vf := NewVentField(5, 5)
	pair := types.PointPair{{2, 1}, {2, 4}}
	vf.AddLine(pair)
	pair = types.PointPair{{1, 2}, {4, 2}}
	vf.AddLine(pair)
	expected := 1
	result := vf.CrossingPointCount(2)
	if expected != result {
		t.Fatalf("Expected: %d, got: %d\n", expected, result)
	}
}

func TestCrossingPointCountDiagonal(t *testing.T) {
	vf := NewVentField(5, 5)
	pair := types.PointPair{{2, 1}, {2, 4}}
	vf.AddLine(pair)
	pair = types.PointPair{{1, 2}, {4, 2}}
	vf.AddLine(pair)
	pair = types.PointPair{{1, 0}, {3, 2}}
	vf.AddLine(pair)
	expected := 3
	result := vf.CrossingPointCount(2)
	if expected != result {
		t.Fatalf("Expected: %d, got: %d\n", expected, result)
	}
}

func TestCalculatePointsToAdd(t *testing.T) {
	pair := types.PointPair{
		{X: 1, Y: 5},
		{X: 3, Y: 5},
	}
	expected := []types.Point{{1, 5}, {2, 5}, {3, 5}}
	result := calculatePointsToAdd(pair)
	if len(expected) != len(result) {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
	for i := 0; i < len(result); i++ {
		if expected[i].X != result[i].X || expected[i].Y != result[i].Y {
			t.Fatalf("Expected: %v, got: %v", expected, result)
		}
	}
	pair = types.PointPair{
		{X: 1, Y: 1},
		{X: 1, Y: 3},
	}
	expected = []types.Point{{1, 1}, {1, 2}, {1, 3}}
	result = calculatePointsToAdd(pair)
	if len(expected) != len(result) {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
	for i := 0; i < len(result); i++ {
		if expected[i].X != result[i].X || expected[i].Y != result[i].Y {
			t.Fatalf("Expected: %v, got: %v", expected, result)
		}
	}
}

func TestCalculatePointsToAddDiagonal(t *testing.T) {
	pair := types.PointPair{
		{X: 1, Y: 1},
		{X: 3, Y: 3},
	}
	expected := []types.Point{{1, 1}, {2, 2}, {3, 3}}
	result := calculatePointsToAdd(pair)
	if len(expected) != len(result) {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
	for i := 0; i < len(result); i++ {
		if expected[i].X != result[i].X || expected[i].Y != result[i].Y {
			t.Fatalf("Expected: %v, got: %v", expected, result)
		}
	}
	pair = types.PointPair{
		{X: 3, Y: 2},
		{X: 1, Y: 4},
	}
	expected = []types.Point{{1, 4}, {2, 3}, {3, 2}}
	result = calculatePointsToAdd(pair)
	if len(expected) != len(result) {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
	for i := 0; i < len(result); i++ {
		if expected[i].X != result[i].X || expected[i].Y != result[i].Y {
			t.Fatalf("Expected: %v, got: %v", expected, result)
		}
	}
}
