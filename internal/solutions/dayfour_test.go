package solutions

import (
	"github.com/etrexel/advent-of-code-2021/pkg/bingocard"
	"testing"
)

var bingoNumbers = []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
var bingoCardInput = []string{
	"" +
		"22 13 17 11  0\n" +
		" 8  2 23  4 24\n" +
		"21  9 14 16  7\n" +
		" 6 10  3 18  5\n" +
		" 1 12 20 15 19",

	"" +
		" 3 15  0  2 22\n" +
		" 9 18 13 17  5\n" +
		"19  8  7 25 23\n" +
		"20 11 10 24  4\n" +
		"14 21 16 12  6",

	"" +
		"14 21 17 24  4\n" +
		"10 16 15  9 19\n" +
		"18  8 23 26 20\n" +
		"22 11 13  6  5\n" +
		" 2  0 12  3  7",
}

func TestDayFourPartOne(t *testing.T) {
	bingoCard, err := bingocard.NewBingoCard(bingoCardInput[2])
	if err != nil {
		t.Fatalf("Error building bingo card\n")
	}
	expected := 4512
	result := DayFourPartOne(bingoNumbers, []bingocard.BingoCard{bingoCard})
	if result != expected {
		t.Fatalf("Expected: %d, got: %d", expected, result)
	}
}

func TestDayFourPartTwo(t *testing.T) {
	bingoCards := make([]bingocard.BingoCard, len(bingoCardInput))
	for i, input := range bingoCardInput {
		bingoCard, err := bingocard.NewBingoCard(input)
		if err != nil {
			t.Fatalf("Error building bingo card\n")
		}
		bingoCards[i] = bingoCard
	}
	expected := 1924
	result := DayFourPartTwo(bingoNumbers, bingoCards)
	if result != expected {
		t.Fatalf("Expected: %d, got: %d", expected, result)
	}
}
