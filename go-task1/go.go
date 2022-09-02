package main

import (
	"fmt"
	"os"
)

func main() {
	var number int
	var amount_of_number int
	fmt.Println("Input n")
	fmt.Scanf("%d", &amount_of_number)
	var array []int
	fmt.Println("Input array")
	for i := 0; i < amount_of_number; i++ {
		fmt.Fscan(os.Stdin, &number)
		array = append(array, number)
	}
	fizzbuzz(array)
}

func fizzbuzz(number []int) {
	fmt.Println("Result:")
	for _, numb := range number {
		switch {
		case numb == 0:
			fmt.Println(numb)
		case numb%3 == 0 && numb%5 == 0:
			fmt.Println("FizzBuzz")
		case numb%3 == 0:
			fmt.Println("Fizz")
		case numb%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(numb)
		}
	}
}
