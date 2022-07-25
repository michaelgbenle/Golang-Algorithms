package main

import (
	"strconv"
	"strings"
)

//Create a function that determines whether a number is Oddish or
//Evenish.A number is Oddish if the sum of all of its digits is odd,
//and a number is Evenish if the sum of all of its digits is even.If
//a number is Oddish, return "Oddish". Otherwise, return "Evenish".



func oddishOrEvenish(num int) string {
	converted := strconv.Itoa(num)
	newArr := strings.Split(converted, "")
	var total int
	for _, v := range newArr {
		b, _ := strconv.Atoi(v)
		total += b
	}
	if total%2 == 0 {
		return "Evenish"
	}
	return "oddish"
}
