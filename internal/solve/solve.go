package solve

import (
	"fmt"

	"github.com/etrexel/advent-of-code-2021/internal/solutions"
)

func Run(c *ConfigOptions) error {
	switch c.Day {
	case 1:
		input, err := solutions.DayOneReadInput(c.InputPath + "day_one.txt")
		if err != nil {
			return err
		}
		result := solutions.DayOnePartOne(input)
		fmt.Printf("Day One Part One: %d\n", result)
		result = solutions.DayOnePartTwo(input)
		fmt.Printf("Day One Part Two: %d\n", result)
		break
	case 2:
		input, err := solutions.DayTwoReadInput(c.InputPath + "day_two.txt")
		if err != nil {
			return err
		}
		result, err := solutions.DayTwoPartOne(input)
		if err != nil {
			return err
		}
		fmt.Printf("Day Two Part One: %d\n", result)
		result, err = solutions.DayTwoPartTwo(input)
		if err != nil {
			return err
		}
		fmt.Printf("Day Two Part Two: %d\n", result)
		break
	case 3:
		input, err := solutions.DayThreeReadInput(c.InputPath + "day_three.txt")
		if err != nil {
			return err
		}
		result := solutions.DayThreePartOne(input)
		fmt.Printf("Day Three Part One: %d\n", result)
		result = solutions.DayThreePartTwo(input)
		fmt.Printf("Day Three Part Two: %d\n", result)
		break
	case 4:
		bingoNumbers, bingoCards, err := solutions.DayFourReadInput(c.InputPath + "day_four.txt")
		if err != nil {
			return err
		}
		result := solutions.DayFourPartOne(bingoNumbers, bingoCards)
		fmt.Printf("Day Four Part One: %d\n", result)
		result = solutions.DayFourPartTwo(bingoNumbers, bingoCards)
		fmt.Printf("Day Four Part Two: %d\n", result)
		break
	default:
		fmt.Println("Invalid day")
	}

	return nil
}
