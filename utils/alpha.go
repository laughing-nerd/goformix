package utils

import (
	"strconv"
)

func ValidateString(args []string, value string) string {

	var minLength int
	var maxLength int

	l := len(args)

	if l > 0 {
		minLength, _ = strconv.Atoi(args[0])
	}
	if l > 1 {
		maxLength, _ = strconv.Atoi(args[1])
	}
	valueLength := len(value)

	// discard maxlength and only check if the vaueLength > minlength
	if minLength >= maxLength {
		if valueLength < minLength {
			return "String length is less than the minimum length"
		}
	} else {
		if valueLength < minLength || valueLength > maxLength {
			return "String length is not within the specified range"
		}
	}
	return ""
}
