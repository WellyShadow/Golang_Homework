package checkerstring

// I'm creating a simple package cheakerstring
// It involves:
// check if string not empty
// trim string
// split the string
// string to date

// importing fmt package for basic
// printing & scan operations
var Test = "text"

type Checkstruct struct {
	Strtocheck string
}

func (str *Checkstruct) Notempty() string {
	if str.Strtocheck == "" {
		str.Strtocheck = "Empty String"
	}
	return str.Strtocheck
}
