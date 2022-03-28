package main

import "fmt"

//https://www.hackerrank.com/contests/week-6-monday/challenges/array-left-rotation
func main() {
	fmt.Println(rotateLeft(5, []int32{1, 2, 3, 4, 5}))
}

func rotateLeft(d int32, arr []int32) []int32 {
	a := arr[:d]
	b := arr[d:]
	b = append(b, a...)
	return b
}
