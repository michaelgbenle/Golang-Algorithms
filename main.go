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
	fmt.Println(sevenBoom([]int{1, 2, 3, 4, 5, 6, 7}))
	fmt.Println(sevenBoom([]int{2, 6, 7, 9, 3}))
	fmt.Println(trackTheRobot([]int{10, -10, -10, 10}))
	fmt.Println(trackTheRobot([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(trackTheRobot([]int{1, 0, 2, 0, 3, 0, 4, 0, 5, 0}))
	fmt.Println(upHill(18, 20, 60))
	fmt.Println(upHill(30, 10, 30))

}


