package main




func leapYear(year int) bool {
	for year%4 == 0 {
		return true
	}
	return false
}
