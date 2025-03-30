package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var x, y, str, col int
	//var sizeOfDesk []int
	var sizeOfDesk [2]int
	for i := 0; i < 2; i++ {
		sizeOfDesk[i], _ = strconv.Atoi(os.Args[i+1])
	}
	//sizeOfDesk = os.Args
	y = sizeOfDesk[0] // количество столбцов
	x = sizeOfDesk[1] // количество строк
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
