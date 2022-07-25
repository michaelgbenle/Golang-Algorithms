package main



//Write a function that takes a number and returns the perimeter
//of either a circle or a square. The input will be in the form (letter l, number num)
//where the letter will be either "s" for square, or "c" for circle, and the number will be the
//side of the square or the radius of the circle.


func perimeter(x string, y float64) float64 {
	if x == "s" {
		return 4 * y
	}
	return 6.28 * y
}
