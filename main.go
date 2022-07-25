package main

import "fmt"

func main() {
	fmt.Println(boomerang([]int{3, 7, 3, 2, 1, 5, 1, 2, 2, -2, 2}))
	fmt.Println(boomerang([]int{2, 2, 2, 2, 3, 2, 3}))
	fmt.Println(backToHome("NENESSWW"))
	fmt.Println(backToHome("EEWE"))
	fmt.Println(perimeter("s", 7))
	fmt.Println(perimeter("s", 4))
	fmt.Println(leapYear(2000))
	fmt.Println(leapYear(1521))
	fmt.Println(rotateLeft(5, []int32{1, 2, 3, 4, 5}))

}


