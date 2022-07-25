package main



func boomerang(boo []int) int {
	var count int
	for i := 0; i < len(boo)-2; i++ {
		var k = i + 1
		var j = i + 2
		if boo[i] == boo[j] && boo[k] != boo[i] {
			count++
		}
	}
	return count
}
