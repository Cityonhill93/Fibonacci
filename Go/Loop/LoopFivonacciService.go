package loop

import (
	"Fibonacci/common"
)

type LoopFibonacciService struct {
}

func (s LoopFibonacciService) GetNumbers(count int) []int {
	var numbers []int = nil

	for i := 0; i < count; i++ {
		if i > 0 {
			var x, y int = common.GetPriorValues(numbers)
			var nextValue = common.GetNextValue(x, y)
			numbers = append(numbers, nextValue)
		} else {
			numbers = append(numbers, 0)
		}
	}

	return numbers
}
