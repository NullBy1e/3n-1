package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	fmt.Println("MainNumber, Parts")
	numbers_raw := os.Args[1]
	numbers, err := strconv.Atoi(numbers_raw)
	if err != nil {
		panic(err)
	}
	results := make(chan string, numbers)
	var wg sync.WaitGroup
	for i := 1; i <= numbers; i++ {
		wg.Add(1)
		go func(){
			defer wg.Done()
			result_raw := calculateCollatz(i)
			result := strings.Join(result_raw,"")
			results <- result
		}()
		fmt.Println(calculateCollatz(i))
	}
	go func() {
		wg.Wait()
		close(results)
	}()
}

func calculateCollatz(number int) []string {
	var result []string
	var numbers string
	result = append(result, strconv.Itoa(number)+",")
	if number == 1 || number == 2 {
		numbers = ""
	} else {
		if number%2 == 0 {
			final_number := number / 2
			numbers = strings.Join(calculateCollatz(final_number), "")
		} else {
			final_number := number*3 + 1
			numbers = strings.Join(calculateCollatz(final_number), "")
		}
	}
	result = append(result, numbers)
	return result
}
