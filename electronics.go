package main

//https://www.hackerrank.com/challenges/electronics-shop/problem

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {

	var tempMax int32

	for i := 0; i < len(keyboards); i++ {
		for j := 0; j < len(drives); j++ {
			if keyboards[i]+drives[j] > tempMax && keyboards[i]+drives[j] <= b {
				tempMax = keyboards[i] + drives[j]
			}
		}

	}
	if tempMax == int32(0) {
		return -1
	}
	return tempMax
}
