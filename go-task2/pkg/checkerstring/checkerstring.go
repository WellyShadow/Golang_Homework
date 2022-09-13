// I'm creating a simple package cheakerstring
// It involves:
// check if string not empty
// trim string
// concat the string
package checkerstring

// The structure contains the string to be processed
type Checkstruct struct {
	Strtocheck string
}

// Method Default checking
// if string empty
// return input variable def
// if string not empty
// return just string
func (str *Checkstruct) Default(def string) string {
	if str.Strtocheck == "" {
		return def
	}
	return str.Strtocheck
}

// Method Trim
// trims the string,
// leaving the number of elements specified by the user
// take string and integer number
// return trimed string
func (str *Checkstruct) Trim(goal int) string {
	rStr := ""
	temp := ""
	for i, srune := range str.Strtocheck {
		temp = string(srune)
		rStr += temp
		if i+1 == goal {
			break
		}
	}
	return rStr
}

// Function Concat
// concatenates strings
// the first parameter is the connector string
// all other parameters, strings to be combined
// return combined string
func Concat(format string, args ...string) string {
	cstr := ""
	for i, arg := range args {
		if i == 0 {
			cstr += arg
		} else if i > 0 {
			cstr += format + arg
		}
	}
	return cstr
}
