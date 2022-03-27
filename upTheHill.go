package main

import "fmt"

//Write a function that returns the average speed traveled given an uphill time,
//uphill rate and a downhill rate. Uphill time is given in minutes. Return the
//rate as an integer (mph). No rounding is necessary.

func main() {
	fmt.Println(upHill(18, 20, 60))
	fmt.Println(upHill(30, 10, 30))
	fmt.Println(upHill(18, 20, 60))
	fmt.Println(upHill(30, 8, 24))
	fmt.Println(upHill(18, 10, 30))
}

func upHill(a, b, c float64) float64 {
	t1 := a / 60
	distance := b * t1
	Tdistance := 2 * distance
	t2 := distance / c
	Ttime := t1 + t2
	avgSpeed := Tdistance / Ttime

	return avgSpeed
}
