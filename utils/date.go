package utils

import (
	"reflect"
)

func ValidateDate(val interface{}) string {
	if !(reflect.TypeOf(val).String() == "time.Time") {
		return "Invalid date value"
	}
	return ""
}
