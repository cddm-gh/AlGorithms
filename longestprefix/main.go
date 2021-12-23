package main

import (
	"fmt"
	s "strings"
)

func main() {

	strs := []string{"hola", "holanda", "hollada"}
	result := longestCommonPrefix(strs)
	fmt.Printf("The longest prefix is: %s", result)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i <= len(strs)-1; i++ {
		for !s.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}
