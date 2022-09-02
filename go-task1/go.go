package main

import (
	"fmt"
	"os"
	"strconv"
)

func Error(err error) {
	if err != nil {
		fmt.Println("Your input incorrect. You must input unsigned integer numbers!")
		os.Exit(0)
	}
}

func main() {
	var number uint64
	var amount_of_number uint64
	fmt.Println("Input amount of numbers")
	_, err := fmt.Scanf("%d", &amount_of_number)
	Error(err)
	var array []uint64
	fmt.Println("Input array")
	for i := uint64(0); i < amount_of_number; i++ {
		_, err := fmt.Fscan(os.Stdin, &number)
		Error(err)
		array = append(array, number)
	}
	fmt.Println(fizzbuzz(array))

}

func fizzbuzz(number []uint64) []string {
	var str []string
	fmt.Println("Result:")
	for _, numb := range number {
		switch {
		case numb == 0:
			str = append(str, strconv.FormatUint(numb, 10))
		case numb%3 == 0 && numb%5 == 0:
			str = append(str, "FizzBuzz")
		case numb%3 == 0:
			str = append(str, "Fizz")
		case numb%5 == 0:
			str = append(str, "Buzz")
		default:
			str = append(str, strconv.FormatUint(numb, 10))
		}
	}
	return str
}
