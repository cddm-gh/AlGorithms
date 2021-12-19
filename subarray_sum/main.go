package main

import "fmt"

func main() {
	var N int // elements in the Array
	_, err := fmt.Scanf("%d", &N)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	A := make([]int, N)
	s := make([]int, N)
	sum := 0
	for i := 0; i < N; i++ {
		_, err := fmt.Scanf("%d", &A[i])
		if err != nil {
			fmt.Println("Error reading A: ", err.Error())
		}
		sum += A[i]
		s[i] = sum
	}

	var C int // number of cases to evaluate
	_, errC := fmt.Scanf("%d", &C)
	if errC != nil {
		fmt.Println("Error reading C: ", err.Error())
	}

	for k := 0; k < C; k++ {
		var p, q int
		_, err := fmt.Scanf("%d %d", &p, &q)
		if err != nil {
			fmt.Println("Error reading p and q: ", err.Error())
		}
		biggestSum := solve(s, p, q)
		fmt.Printf("%d\n", biggestSum)
	}
}

func solve(s []int, p, q int) int {
	sum := s[q]
	if p-1 >= 0 {
		sum -= s[p-1]
	}
	return sum
}
