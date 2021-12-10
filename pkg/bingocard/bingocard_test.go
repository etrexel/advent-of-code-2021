package bingocard

import "testing"

var testInput = "" +
	"14 21 17 24  4\n" +
	"10 16 15  9 19\n" +
	"18  8 23 26 20\n" +
	"22 11 13  6  5\n" +
	" 2  0 12  3  7"

func TestMarkNumber(t *testing.T) {
	card, err := NewBingoCard(testInput)
	if err != nil {
		t.Fatalf("Could not create bingo card: %s\n", err)
	}
	card.MarkNumber(21)
	if card.marked[21] != true {
		t.Fatalf("Expected: true, got: false\n")
	}
}

func TestHasWon(t *testing.T) {
	card, err := NewBingoCard(testInput)
	if err != nil {
		t.Fatalf("Could not create bingo card: %s\n", err)
	}
	card.MarkNumber(21)
	card.MarkNumber(14)
	if card.HasWon() != false {
		t.Fatalf("Expected: false, got: true\n")
	}
	card.MarkNumber(17)
	card.MarkNumber(4)
	card.MarkNumber(24)
	if card.HasWon() != true {
		t.Fatalf("Expected: true, got: false\n")
	}
}

func TestWinningSum(t *testing.T) {
	card, err := NewBingoCard(testInput)
	if err != nil {
		t.Fatalf("Could not create bingo card: %s\n", err)
	}
	input := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24}
	for _, val := range input {
		card.MarkNumber(val)
	}
	expected := 188
	result := card.WinningSum()
	if result != expected {
		t.Fatalf("Expected: %d, got: %d\n", expected, result)
	}
}

func TestCheckRow(t *testing.T) {
	card, err := NewBingoCard(testInput)
	if err != nil {
		t.Fatalf("Could not create bingo card: %s\n", err)
	}
	card.MarkNumber(18)
	card.MarkNumber(8)
	card.MarkNumber(23)
	if card.checkRow(2) != false {
		t.Fatalf("Expected: false, got: true\n")
	}
	card.MarkNumber(26)
	card.MarkNumber(20)
	if card.checkRow(2) != true {
		t.Fatalf("Expected: true, got: false\n")
	}
}

func TestCheckColumn(t *testing.T) {
	card, err := NewBingoCard(testInput)
	if err != nil {
		t.Fatalf("Could not create bingo card: %s\n", err)
	}
	card.MarkNumber(24)
	card.MarkNumber(9)
	card.MarkNumber(26)
	if card.checkColumn(3) != false {
		t.Fatalf("Expected: false, got: true\n")
	}
	card.MarkNumber(6)
	card.MarkNumber(3)
	if card.checkColumn(3) != true {
		t.Fatalf("Expected: true, got: false\n")
	}
}
