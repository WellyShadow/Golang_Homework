package checkerstring

// I'm creating a simple package cheakerstring
// It involves:
// check if string not empty
// trim string
// split the string
// string to date

//The structure contains the string to be processed
type Checkstruct struct {
	Strtocheck string
}

// Method Notempty checking
// if string empty
// return "Empty String" in variable
// if string not empty
// return just string
func (str *Checkstruct) Notempty() string {
	if str.Strtocheck == "" {
		str.Strtocheck = "Empty String"
	}
	return str.Strtocheck
}

//Function Trim
// trims the string,
// leaving the number of elements specified by the user
// take string and integer number
// return trimed string
func Trim(str Checkstruct, goal int) string {
	return_str := ""
	temp := ""
	for index, srune := range str.Strtocheck {
		temp = string(srune)
		return_str += temp
		if index+1 == goal {
			break
		}
	}
	return return_str
}
