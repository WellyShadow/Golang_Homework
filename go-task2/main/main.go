package main

import (
	"fmt"

	"github.com/WellyShadow/Golang_Homework/go-task2/pkg/checkerstring"
)

func main() {
	s := "Vlad"
	fmt.Println("Hi")
	checkerstring.Notempty(s)
	fmt.Println(s)
}
