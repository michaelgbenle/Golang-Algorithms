package main



//https://www.hackerrank.com/contests/week-6-monday/challenges/array-left-rotation


func rotateLeft(d int32, arr []int32) []int32 {
	a := arr[:d]
	b := arr[d:]
	b = append(b, a...)
	return b
}
