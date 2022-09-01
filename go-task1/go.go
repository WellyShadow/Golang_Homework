package main

import (
	"fmt"
	"os"
)

func main() {
	var number int
	var n int
	fmt.Printf("Input n")
	fmt.Scanf("%d", &n)

	var array []int
	fmt.Println("Input array")
	fmt.Println(array)
	for i := 0; i < n; i++ {

		fmt.Fscan(os.Stdin, &number)
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
