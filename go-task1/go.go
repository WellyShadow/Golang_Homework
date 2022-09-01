package main

import "fmt"

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fizzbuzz(input)

}
func fizzbuzz(number []int) {
	for i := 0; i < len(number); i++ {
		if number[i]%3 == 0 && number[i]%5 == 0 {
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
