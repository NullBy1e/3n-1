package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("MainNumber, Parts")
	numbers_raw := os.Args[1]
	numbers, err := strconv.Atoi(numbers_raw)
	if err != nil {
		panic(err)
	}
	for i := 1; i <= numbers; i++ {
		calculateCollatz(i)
	}
}

func calculateCollatz(number int) {
	fmt.Print(strconv.Itoa(number) + ",")
	if number == 1 || number == 2 {
		fmt.Print("\n")
	} else {
		if number%2 == 0 {
			final_number := number / 2
			calculateCollatz(final_number)
		} else {
			final_number := number*3 + 1
			calculateCollatz(final_number)
		}
	}
}
