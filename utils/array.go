package utils

import (
	"reflect"
)

func ValidateArray(intendedType reflect.Type, val interface{}) string {
	if !(reflect.TypeOf(val) == intendedType) {
		return "Invalid array"
	}
  return ""
}
