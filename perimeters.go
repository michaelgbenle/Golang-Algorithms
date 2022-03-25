package main

import "fmt"

//Write a function that takes a number and returns the perimeter
//of either a circle or a square. The input will be in the form (letter l, number num)
//where the letter will be either "s" for square, or "c" for circle, and the number will be the
//side of the square or the radius of the circle.

func main() {
	fmt.Println(perimeter("s", 7))
	fmt.Println(perimeter("s", 4))
	fmt.Println(perimeter("c", 1))
	fmt.Println(perimeter("c", 4))
	fmt.Println(perimeter("c", 9))
}

func perimeter(x string, y float64) float64 {
	if x == "s" {
		return 4 * y
	}
	return 6.28 * y
}
