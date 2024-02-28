package main

import (
	"Fibonacci/common"
)

type RecursiveFibonacciService struct {
}

func (s RecursiveFibonacciService) GetNumbers(count int) []int {
	return getNumbersRecursive(nil, count)
}

func getNumbersRecursive(numbers []int, count int) []int {
	var outNumbers []int

	if len(numbers) > 0 {
		var x, y int = common.GetPriorValues(numbers)
		var nextValue int = common.GetNextValue(x, y)
		outNumbers = append(numbers, nextValue)
	} else {
		outNumbers = append(numbers, 0)
	}

	if len(outNumbers) < count {
		outNumbers = getNumbersRecursive(outNumbers, count)
	}

	return outNumbers
}
