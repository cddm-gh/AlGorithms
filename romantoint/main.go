package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	result := romanToInt("MCMXCIV")
	elapse := time.Since(start)
	fmt.Printf("find the int value %d took %s\n", result, elapse)
}

func romanToInt(s string) int {
	var value, lastValue int
	romanLiterals := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := len(s) - 1; i >= 0; i-- {
		currentValue := romanLiterals[s[i]]
		if currentValue < lastValue {
			value -= currentValue
		} else {
			value += currentValue
		}
		lastValue = currentValue
	}

	return value
}
