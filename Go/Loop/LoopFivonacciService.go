package Loop

import (
	"Fibonacci/Common"
)

type LoopFibonacciService struct {
}

func (s LoopFibonacciService) GetNumbers(count int) []int {
	var numbers []int = nil

	for i := 0; i < count; i++ {
		if i > 0 {
			var x, y int = Common.GetPriorValues(numbers)
			var nextValue = Common.GetNextValue(x, y)
			numbers = append(numbers, nextValue)
		} else {
			numbers = append(numbers, 0)
		}
	}

	return numbers
}
