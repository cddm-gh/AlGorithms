package main

import "fmt"

func main() {
	numbers := []int{-5, 2, 7, 9, 11, 15}
	target := 26

	result := solve(numbers, target)
	fmt.Printf("The result is: %v\n", result)
}

func solve(numbers []int, target int) []int {
	lp := 0
	rp := len(numbers) - 1
	idx := make([]int, 2)
	for {
		total := numbers[lp] + numbers[rp]
		if total == target {
			idx[0] = lp + 1
			idx[1] = rp + 1
			break
		} else if total > target {
			// decrement the right pointer
			rp -= 1
		} else {
			// increment the left pointer
			lp += 1
		}
	}
	return idx
}
