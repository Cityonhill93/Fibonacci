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

	if valueOne != 3 || valueTwo != 5 {
		t.Fail()
	}
}
