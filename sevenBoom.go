package main

import "fmt"

//Create a function that takes an array of numbers and return "Boom!"
//if the digit 7 appears in the array. Otherwise, return "there is no 7 in the array".

func main() {
	fmt.Println(sevenBoom([]int{1, 2, 3, 4, 5, 6, 7}))
	fmt.Println(sevenBoom([]int{2, 6, 7, 9, 3}))
	fmt.Println(sevenBoom([]int{33, 68, 400, 5}))
	fmt.Println(sevenBoom([]int{186, 48, 100, 66}))
	fmt.Println(sevenBoom([]int{76, 55, 44, 32}))
}

func sevenBoom(arr []int) string {
	for _, v := range arr {
		if v == 7 {
			return "Boom!"
		}
	}
	return "there is no 7 in the array"
}
