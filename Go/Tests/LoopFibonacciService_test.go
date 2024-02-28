package Tests

import (
	"Fibonacci/Common"
	"Fibonacci/Loop"
	"testing"
)

func TestGetNumbersLoop(t *testing.T) {
	var service Common.IFibonacciService = Loop.LoopFibonacciService{}

	ServiceReturnsCorrectNumbers(service, t)
}
