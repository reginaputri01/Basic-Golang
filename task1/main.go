package main

import (
	"fmt"
)

func cetakGambar(length int) {
	value := ""
	center := (length / 2) + 1

	for a := 1; a <= length; a++ {
		for b := 1; b <= length; b++ {
			if a == center {
				value += "*"
				value += " "
			} else {
				if b == 1 || b == length {
					value += "*"
					value += " "
				} else {
					value += "="
					value += " "
				}
			}
		}
		fmt.Println(value)
		value = ""
	}
}

func main() {
	cetakGambar(5)
}
