package ventfield

import (
	"math"

	"github.com/etrexel/advent-of-code-2021/pkg/types"
)

type VentField struct {
	field [][]int
}

func NewVentField(xDim int, yDim int) VentField {
	vf := VentField{field: make([][]int, yDim)}
	for i, _ := range vf.field {
		vf.field[i] = make([]int, xDim)
	}
	return vf
}

func (vf *VentField) AddLine(pointPair types.PointPair) {
	points := calculatePointsToAdd(pointPair)
	for _, point := range points {
		vf.field[point.Y][point.X]++
	}
}

func (vf *VentField) CrossingPointCount(count int) int {
	total := 0
	for _, row := range vf.field {
		for _, val := range row {
			if val >= count {
				total++
			}
		}
	}

	return total
}

func calculatePointsToAdd(pointPair types.PointPair) []types.Point {
	xOffset := math.Abs(float64(pointPair[0].X) - float64(pointPair[1].X))
	yOffset := math.Abs(float64(pointPair[0].Y) - float64(pointPair[1].Y))
	var output []types.Point
	if xOffset == 0 {
		var startingPoint types.Point
		var endingPoint types.Point
		if pointPair[0].Y > pointPair[1].Y {
			startingPoint = pointPair[1]
			endingPoint = pointPair[0]
		} else {
			startingPoint = pointPair[0]
			endingPoint = pointPair[1]
		}
		for i := startingPoint.Y; i <= endingPoint.Y; i++ {
			output = append(output, types.Point{X: startingPoint.X, Y: i})
		}
	} else {
		var startingPoint types.Point
		var endingPoint types.Point
		if pointPair[0].X > pointPair[1].X {
			startingPoint = pointPair[1]
			endingPoint = pointPair[0]
		} else {
			startingPoint = pointPair[0]
			endingPoint = pointPair[1]
		}
		var slope int
		if yOffset == 0 {
			slope = 0
		} else if startingPoint.Y < endingPoint.Y {
			slope = 1
		} else {
			slope = -1
		}
		intercept := -(slope*startingPoint.X - startingPoint.Y)
		for i := startingPoint.X; i <= endingPoint.X; i++ {
			point := types.Point{X: i, Y: slope*i + intercept}
			output = append(output, point)
		}
	}
	return output
}
