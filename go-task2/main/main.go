package main

import (
	"fmt"

	"github.com/WellyShadow/Golang_Homework/go-task2/pkg/checkerstring"
)

func main() {
	s := checkerstring.Checkstruct{""}
	s.Notempty()
	fmt.Println(s)

	strtortim := checkerstring.Checkstruct{"123456789"}
	straftertrim := checkerstring.Trim(strtortim, 6)
	fmt.Println(straftertrim)
}
