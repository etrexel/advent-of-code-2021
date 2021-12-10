package solutions

import (
	"github.com/etrexel/advent-of-code-2021/pkg/bingocard"
	"os"
	"sort"
	"strconv"
	"strings"
)

func DayFourReadInput(filePath string) ([]int, []bingocard.BingoCard, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, nil, err
	}
	values := strings.Split(string(data), "\n\n")
	bingoNumbers := make([]int, len(values[0]))
	for idx, value := range strings.Split(values[0], ",") {
		i, err := strconv.Atoi(value)
		if err != nil {
			return nil, nil, err
		}
		bingoNumbers[idx] = i
	}

	bingoCards := make([]bingocard.BingoCard, len(values[1:]))
	for idx, value := range values[1:] {
		bingoCard, err := bingocard.NewBingoCard(value)
		if err != nil {
			return nil, nil, err
		}
		bingoCards[idx] = bingoCard
	}
	return bingoNumbers, bingoCards, nil
}

func DayFourPartOne(bingoNumbers []int, bingoCards []bingocard.BingoCard) int {
	for _, number := range bingoNumbers {
		for _, card := range bingoCards {
			card.MarkNumber(number)
			if card.HasWon() {
				return card.WinningSum() * number
			}
		}
	}
	return 0
}

func DayFourPartTwo(bingoNumbers []int, bingoCards []bingocard.BingoCard) int {
	for _, number := range bingoNumbers {
		var markForDelete []int
		for i, card := range bingoCards {
			card.MarkNumber(number)
			if card.HasWon() && len(bingoCards) == 1 {
				return card.WinningSum() * number
			} else if card.HasWon() {
				markForDelete = append(markForDelete, i)
			}
		}
		sort.Ints(markForDelete)
		for i := len(markForDelete) - 1; i >= 0; i-- {
			marker := markForDelete[i]
			bingoCards = append(bingoCards[:marker], bingoCards[marker+1:]...)
		}
	}
	return 0
}
