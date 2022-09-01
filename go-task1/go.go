package main

import (
	"fmt"
)

func main() {
	number := 0
	array := make([]int, 0)
	fmt.Println("Input array")
	for i := 0; i < 3; i++ {
		fmt.Scanf("%d", &number)
		array = append(array, number)
	}
	fmt.Println(array)
	fizzbuzz(array)

}

func fizzbuzz(number []int) {
	for i := 0; i < len(number); i++ {
		if number[i] == 0 {
			fmt.Println(number[i])
		} else if number[i]%3 == 0 && number[i]%5 == 0 {
			fmt.Println(number[i], "FizzBuzz")
		} else if number[i]%3 == 0 {
			fmt.Println(number[i], "Fizz")
		} else if number[i]%5 == 0 {
			fmt.Println(number[i], "Buzz")
		} else {
			fmt.Println(number[i])
		}
	}
}
