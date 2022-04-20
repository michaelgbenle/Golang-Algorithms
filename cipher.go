package main

import "fmt"

func main() {
	//fmt.Println(cipher("middle-Outz", 2))
	//fmt.Println(cipher("A friend in need is a friend indeed", 20))
	//fmt.Println(cipher("Always-Look-on-the-Bright-Side-of-Life", 5))
	//fmt.Println(cipher("A Fool and His Money Are Soon Parted.", 27))
	fmt.Println(cipher("One should not worry over things that have already happened and that cannot be changed.", 49))
	//"Always-Look-on-the-Bright-Side-of-Life"
	//"A friend in need is a friend indeed"
	//"U zlcyhx ch hyyx cm u zlcyhx chxyyx"
}

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
