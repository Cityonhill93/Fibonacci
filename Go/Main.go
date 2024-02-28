package main

import (
	"Fibonacci/common"
	"Fibonacci/loop"
	"Fibonacci/recursion"
	"fmt"
)

func main() {
	const digitCount int = 10

	fmt.Println("Welcome to the Fibonacci app!")

	fmt.Println("Running the loop service...")
	printNumbers(digitCount, loop.LoopFibonacciService{})

	fmt.Println("Running the recursive service...")
	printNumbers(digitCount, recursion.RecursiveFibonacciService{})

	var discard string
	fmt.Scanln(discard)
}

func printNumbers(count int, service common.IFibonacciService) {
	var numbers []int = service.GetNumbers(count)

	fmt.Println("Start...")
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	fmt.Println("End...")
}
