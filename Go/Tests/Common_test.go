package Tests

import (
	"Fibonacci/Common"
	"testing"
)

func TestGetNextValueReturnsCorrectValue(t *testing.T) {
	var nextValue = Common.GetNextValue(1, 1)

	if nextValue != 2 {
		t.Fail()
	}
}

func TestGetPriorValuesReturnsTwoLatestValues(t *testing.T) {
	var allValues [6]int = [6]int{0, 1, 1, 2, 3, 5}
	var valueOne, valueTwo int = Common.GetPriorValues(allValues[:])

	if valueOne != 5 || valueTwo != 3 {
		t.Fail()
	}
}

func ServiceReturnsCorrectNumbers(service Common.IFibonacciService, t *testing.T) {
	var numbers = service.GetNumbers(5)

	if numbers[0] != 0 || numbers[1] != 1 || numbers[2] != 1 || numbers[3] != 2 || numbers[4] != 3 {
		t.Fail()
	}
}
