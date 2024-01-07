package main

func getNextValue(x, y int) int {
	var sums int = x + y
	if sums == 0 {
		return 1
	}

	return sums
}

func getPriorValues(numbers []int) (int, int) {
	var maxIndex int = len(numbers) - 1

	var priorValue int = getValueOrDefault(maxIndex, numbers)
	var twoValuesAgo int = getValueOrDefault(maxIndex-1, numbers)

	return priorValue, twoValuesAgo
}

func getValueOrDefault(index int, numbers []int) int {
	if index >= 0 {
		return numbers[index]
	} else {
		return 0
	}
}
