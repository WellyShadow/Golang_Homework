package main

import (
	"fmt"

	"github.com/WellyShadow/Golang_Homework/go-task2/pkg/checkerstring"
)

func main() {
	s := checkerstring.Checkstruct{"123"}
	s.Default("Empty")
	fmt.Println(s)

	strtotrim := checkerstring.Checkstruct{"123456789"}
	result := strtotrim.Trim(3)
	fmt.Println(result)

	cstr1 := "12"
	cstr2 := "34"
	cstr3 := "56"
	cstr := checkerstring.Concat(",", cstr1, cstr2, cstr3)
	fmt.Println(cstr)
}
