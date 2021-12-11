package solutions

import (
	"os"
	"strconv"
	"strings"

	"github.com/etrexel/advent-of-code-2021/pkg/types"
	"github.com/etrexel/advent-of-code-2021/pkg/ventfield"
)

func DayFiveReadInput(filePath string) ([]types.PointPair, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return []types.PointPair{}, err
	}
	values := strings.Split(string(data), "\n")
	output := make([]types.PointPair, len(values))
	for i, line := range values {
		pairs := strings.Split(line, " -> ")
		var pointPair types.PointPair
		for j, pointStr := range pairs {
			coords := strings.Split(pointStr, ",")
			x, err := strconv.Atoi(coords[0])
			if err != nil {
				return []types.PointPair{}, err
			}
			y, err := strconv.Atoi(coords[1])
			if err != nil {
				return []types.PointPair{}, err
			}
			pointPair[j] = types.Point{
				X: x,
				Y: y,
			}
		}
		output[i] = pointPair
	}
	return output, nil
}

func DayFivePartOne(points []types.PointPair) int {
	x, y := calculateVentFieldDimensions(points)
	ventField := ventfield.NewVentField(x, y)
	for _, pair := range points {
		if pair[0].X != pair[1].X && pair[0].Y != pair[1].Y {
			continue
		}
		ventField.AddLine(pair)
	}
	return ventField.CrossingPointCount(2)
}

func DayFivePartTwo(points []types.PointPair) int {
	x, y := calculateVentFieldDimensions(points)
	ventField := ventfield.NewVentField(x, y)
	for _, pair := range points {
		ventField.AddLine(pair)
	}
	return ventField.CrossingPointCount(2)
}

func calculateVentFieldDimensions(points []types.PointPair) (int, int) {
	maxX, maxY := 0, 0
	for _, pair := range points {
		if pair[0].X > maxX {
			maxX = pair[0].X
		}
		if pair[1].X > maxX {
			maxX = pair[1].X
		}
		if pair[0].Y > maxY {
			maxY = pair[0].Y
		}
		if pair[1].Y > maxY {
			maxY = pair[1].Y
		}
	}
	return maxX + 1, maxY + 1
}
