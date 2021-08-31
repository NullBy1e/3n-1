package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("MainNumber, Parts")
	for i := 1; i > 0; i++ {
		calculateCollatz(i)
	}
}

func calculateCollatz(number int) {
	fmt.Print(strconv.Itoa(number))
	if number == 1 {
		fmt.Print("\n")
	} else {
		fmt.Print(",")
		if number%2 == 0 {
			final_number := number / 2
			calculateCollatz(final_number)
		} else {
			final_number := number*3 + 1
			calculateCollatz(final_number)
		}
	}
}
