package bingocard

import (
	"strconv"
	"strings"
)

type BingoCard struct {
	card   [][]int
	marked map[int]bool
}

func NewBingoCard(input string) (BingoCard, error) {
	output := BingoCard{
		marked: make(map[int]bool),
	}
	cleanedInput := strings.ReplaceAll(input, "  ", " ")
	cleanedInput = strings.ReplaceAll(cleanedInput, "\n ", "\n")
	if strings.HasPrefix(cleanedInput, " ") {
		cleanedInput = cleanedInput[1:]
	}
	rows := strings.Split(cleanedInput, "\n")
	for rowIdx, row := range rows {
		colValues := strings.Split(row, " ")
		output.card = append(output.card, make([]int, len(colValues)))
		for colIdx, col := range colValues {
			colVal, err := strconv.Atoi(col)
			if err != nil {
				return BingoCard{}, err
			}
			output.card[rowIdx][colIdx] = colVal
		}
	}
	return output, nil
}

func (card *BingoCard) MarkNumber(number int) {
	card.marked[number] = true
}

func (card *BingoCard) HasWon() bool {
	for i, _ := range card.card {
		if card.checkRow(i) || card.checkColumn(i) {
			return true
		}
	}
	return false
}

func (card *BingoCard) WinningSum() int {
	output := 0
	for _, row := range card.card {
		for _, val := range row {
			if !card.marked[val] {
				output += val
			}
		}
	}
	return output
}

func (card *BingoCard) checkRow(row int) bool {
	for _, val := range card.card[row] {
		if card.marked[val] == false {
			return false
		}
	}
	return true
}

func (card *BingoCard) checkColumn(column int) bool {
	for i := 0; i < len(card.card); i++ {
		if card.marked[card.card[i][column]] == false {
			return false
		}
	}
	return true
}
