package main

type IFibonacciService interface {
	GetNumbers(count int) []int
}
