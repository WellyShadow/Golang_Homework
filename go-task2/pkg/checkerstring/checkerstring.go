package checkerstring

// I'm creating a simple package cheakerstring
// It involves:
// check if string not empty
// trim string
// split the string
// string to date

// importing fmt package for basic
// printing & scan operations
var test = "text"

type checkstruct struct {
	strtocheck string
}

func (str checkstruct) Notempty() string {
	if str.strtocheck == "" {
		str.strtocheck = "Empty String"
	}
	return str.strtocheck
}
