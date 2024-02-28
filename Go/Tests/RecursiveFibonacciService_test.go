package Tests

import (
	"Fibonacci/Common"
	"Fibonacci/Recursion"
	"testing"
)

func TestGetNumbersRecursive(t *testing.T) {
	var service Common.IFibonacciService = Recursion.RecursiveFibonacciService{}

	ServiceReturnsCorrectNumbers(service, t)
}
