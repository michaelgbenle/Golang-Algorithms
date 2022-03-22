package main

import "fmt"

func main() {

	result := countTrue([]bool{true, false, false, true, false})
	result2 := countTrue([]bool{false, false, false, false})
	result3 := countTrue([]bool{false, true, false, false})
	fmt.Printf("%v\n", result)
	fmt.Printf("%v\n", result2)
	fmt.Printf("%v\n", result3)
}

func countTrue(items []bool) int {
	var count int
	if len(items) == 0 {
		return 0
	}
	for _, item := range items {
		if item {
			count++
		}
	}
	return count
}
