package main

import "fmt"

func main() {
	x := 0
	result := isPalindrome(x)
	fmt.Printf("%d is palindrome %t\n", x, result)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	var remainder int
	var reversed int64
	check := x

	for check != 0 {
		remainder = check % 10
		reversed = reversed*10 + int64(remainder)
		check /= 10
	}

	return x == int(reversed)
}
