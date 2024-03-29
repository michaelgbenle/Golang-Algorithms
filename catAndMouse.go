package main

import "math"

//https://www.hackerrank.com/challenges/cats-and-a-mouse/problem?isFullScreen=true

func catAndMouse(x int32, y int32, z int32) string {
	po1 := math.Abs(float64(z - x))
	po2 := math.Abs(float64(z - y))
	if po1 < po2 {
		return "Cat A"
	} else if po1 > po2 {
		return "Cat B"
	} else if po1 == po2 {
		return "Mouse C"
	}
	return ""
}
