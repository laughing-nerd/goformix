package utils

import (
	"reflect"
)

func ValidateBoolean(val interface{}) string {
	valType := reflect.TypeOf(val).Kind()

	if !(valType == reflect.Bool) {
		return "Invalid boolean value"
	}
  return ""
}
