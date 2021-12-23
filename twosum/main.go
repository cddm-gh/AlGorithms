package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	var mapVals = make(map[int]int)
	target := 22

	result := solve(nums, target, mapVals)
	fmt.Println(mapVals)
	fmt.Printf("The result is: %v\n", result)
}

func solve(nums []int, target int, mapVals map[int]int) []int {
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		fmt.Printf("target %d val %d, diff: %d\n", target, nums[i], diff)
		if idx, ok := mapVals[diff]; ok {
			fmt.Printf("%d exist on the map\n", diff)
			return []int{idx, i}
		}
		fmt.Printf("%d doesn't exist on the map\n", diff)
		mapVals[nums[i]] = i
		fmt.Printf("Adding %d index %d to map\n", nums[i], i)
	}
	return []int{}
}
