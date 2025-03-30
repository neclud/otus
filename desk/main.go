package main

import "fmt"

func main() {
	var x, y, str, col int
	y = 6 // количество столбцов
	x = 6 // количество строк
	y *= 2
	for str < x {
		for col < y {
			if col%2 == 0 {
				if str%2 == 0 {
					fmt.Print(" ")
				} else {
					fmt.Print("x")
				}
			} else {
				//fmt.Print(" ")
				if str%2 == 0 {
					fmt.Print("x")
				} else {
					fmt.Print(" ")
				}
			}
			col++
		}
		col = 0
		str++
		fmt.Print("\n")
	}
}
