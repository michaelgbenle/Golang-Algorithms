package main



func cipher(s string, n int) string {
	var result string
	if n > 25 {
		n = n % 26
	}
	for _, value := range s {
		if value == 32 || value == 45 || value == 46 {
			result += string(value)
			continue
		}
		if value > 64 && value < 91 {
			if value+int32(n) > 64 && value+int32(n) < 91 {
				result += string(value + int32(n))
			} else {
				result += string(value + int32(n) - int32(26))
			}
		} else {
			if value+int32(n) <= 122 {
				result += string(value + int32(n))
			} else {
				result += string(value + int32(n) - int32(26))
			}
		}

	}
	return result
}
