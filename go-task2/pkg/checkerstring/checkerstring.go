package checkerstring

// I'm creating a simple package cheakerstring
// It involves:
// check if string not empty
// trim string
// split the string
// string to date

// importing fmt package for basic
// printing & scan operations

type checkerstring struct {
	strtocheck string
}

func (str checkerstring) Notempty() string {
	if str.strtocheck == "" {
		str.strtocheck = "Empty String"
	}
	return str.strtocheck
}
