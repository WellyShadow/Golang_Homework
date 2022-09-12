package main

import (
	"fmt"
	
	"github.com/WellyShadow/Golang_Homework/go-task2/pkg/checkerstring"
)

func main() {
	s := checkerstring.checkstruct{"Vlad"}
	fmt.Println("Hi")
	s.Notempty()
	fmt.Println(s)
}
