package utils

import (
	"reflect"
)

func ValidateNumber(val interface{}) string {
	valType := reflect.TypeOf(val).Kind()

	if !(valType == reflect.Int) && !(valType == reflect.Int8) && !(valType == reflect.Int16) && !(valType == reflect.Int32) && !(valType == reflect.Int64) && !(valType == reflect.Uint) && !(valType == reflect.Uint8) && !(valType == reflect.Uint16) && !(valType == reflect.Uint32) && !(valType == reflect.Uint64) && !(valType == reflect.Float32) && !(valType == reflect.Float64) {
		return "Invalid number"
	}
	return ""
}
