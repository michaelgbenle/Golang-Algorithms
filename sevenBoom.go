package main


//Create a function that takes an array of numbers and return "Boom!"
//if the digit 7 appears in the array. Otherwise, return "there is no 7 in the array".



func sevenBoom(arr []int) string {
	for _, v := range arr {
		if v == 7 {
			return "Boom!"
		}
	}
	return "there is no 7 in the array"
}
