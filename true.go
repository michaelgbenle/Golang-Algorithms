package main





func countTrue(items []bool) int {
	var count int
	if len(items) == 0 {
		return 0
	}
	for _, item := range items {
		if item {
			count++
		}
	}
	return count
}
