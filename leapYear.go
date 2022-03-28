package main

import "fmt"

func main() {
	fmt.Println(leapYear(2000))
	fmt.Println(leapYear(1521))
	fmt.Println(leapYear(1996))
	fmt.Println(leapYear(1800))
	fmt.Println(leapYear(2016))
}

func leapYear(year int) bool {
	for year%4 == 0 {
		return true
	}
	return false
}
