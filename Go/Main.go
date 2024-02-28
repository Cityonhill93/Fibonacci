package main

import (
	"Fibonacci/common"
	"fmt"
)

func main() {
	const digitCount int = 10

	fmt.Println("Welcome to the Fibonacci app!")

	fmt.Println("Running the loop service...")
	printNumbers(digitCount, LoopFibonacciService{})

	fmt.Println("Running the recursive service...")
	printNumbers(digitCount, RecursiveFibonacciService{})

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
