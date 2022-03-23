package main

import "fmt"

//BACK TO HOME
//Mubashir has started his journey from home. Given a string of directions
//(N=North, W=West, S=South, E=East), he will walk for one minute in each
//direction.Determine whether a set of directions will lead him back to the starting position or not.

func main() {
	fmt.Println(backToHome("NENESSWW"))
	fmt.Println(backToHome("EEWE"))
	fmt.Println(backToHome("NEESSW"))
	fmt.Println(backToHome("NNSSEEEWWWEW"))
}

func backToHome(direction string) bool {
	vCount := int32(0)
	hCount := int32(0)
	for _, v := range direction {
		if v == 'N' {
			vCount++
		} else if v == 'E' {
			hCount++
		} else if v == 'W' {
			hCount--
		} else if v == 'S' {
			vCount--
		}
	}
	if vCount == 0 && hCount == 0 {
		return true
	}
	return false
}
