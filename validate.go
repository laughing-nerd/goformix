// ----------------------------------------------------------------------------------------------------------------
//
//
//   ,ad8888ba,     ,ad8888ba,    88888888888  ,ad8888ba,    88888888ba   88b           d88  88  8b        d8
//  d8"'    `"8b   d8"'    `"8b   88          d8"'    `"8b   88      "8b  888b         d888  88   Y8,    ,8P
// d8'            d8'        `8b  88         d8'        `8b  88      ,8P  88`8b       d8'88  88    `8b  d8'
// 88             88          88  88aaaaa    88          88  88aaaaaa8P'  88 `8b     d8' 88  88      Y88P
// 88      88888  88          88  88"""""    88          88  88""""88'    88  `8b   d8'  88  88      d88b
// Y8,        88  Y8,        ,8P  88         Y8,        ,8P  88    `8b    88   `8b d8'   88  88    ,8P  Y8,
//  Y8a.    .a88   Y8a.    .a8P   88          Y8a.    .a8P   88     `8b   88    `888'    88  88   d8'    `8b
//   `"Y88888P"     `"Y8888Y"'    88           `"Y8888Y"'    88      `8b  88     `8'     88  88  8P        Y8
//
//
// Version: 0.1.0
// Just another form validation library for Go
//
// ----------------------------------------------------------------------------------------------------------------

// TODO: Proper test cases have to be written
package goformix

import (
	"reflect"
	"strings"
	"sync"

	"github.com/laughing-nerd/goformix/safeMap"
	"github.com/laughing-nerd/goformix/utils"
)

func Validate(formData interface{}) map[string]interface{} {
	message := safemap.New()
	var wg sync.WaitGroup

	t := reflect.TypeOf(formData)
	fields := t.NumField()

	for i := range fields {

		wg.Add(1)
		go func() {
			defer wg.Done()
			field := t.Field(i)
			tag := field.Tag.Get("goformix")

			if len(tag) != 0 {
				tokens := strings.Split(tag, "|")

				switch tokens[0] {
				case "string":
					message.Set(field.Name, utils.ValidateString(tokens[1:], reflect.ValueOf(formData).Field(i).String()))
				case "number":
					message.Set(field.Name, utils.ValidateNumber(reflect.ValueOf(formData).Field(i).Interface()))
				case "boolean":
					message.Set(field.Name, utils.ValidateBoolean(reflect.ValueOf(formData).Field(i).Interface()))
				case "email":
					message.Set(field.Name, utils.ValidateEmail(tokens[1:], reflect.ValueOf(formData).Field(i).String()))
				case "url":
					message.Set(field.Name, utils.ValidateUrl(tokens[1:], reflect.ValueOf(formData).Field(i).String()))
				case "phone":
					message.Set(field.Name, utils.ValidatePhone(tokens[1:], reflect.ValueOf(formData).Field(i).String()))
				case "date":
					message.Set(field.Name, utils.ValidateDate(reflect.ValueOf(formData).Field(i).Interface()))
				case "array":
					message.Set(field.Name, utils.ValidateArray(reflect.ValueOf(formData).Field(i).Type(), reflect.ValueOf(formData).Field(i).Interface()))
				case "embedded":
					message.Set(field.Name, Validate(reflect.ValueOf(formData).Field(i).Interface()))
				case "[]embedded":
					message.Set(field.Name, []map[string]interface{}{})
					arrObj := reflect.ValueOf(formData).Field(i).Interface()
					valSlice := reflect.ValueOf(arrObj)
					for j := 0; j < valSlice.Len(); j++ {

						wg.Add(1)
						go func() {
							defer wg.Done()
							val := valSlice.Index(j).Interface()
							prev, _ := message.Get(field.Name)
							prev = append(prev.([]map[string]interface{}), Validate(val))
							message.Set(field.Name, append(prev.([]map[string]interface{}), Validate(val)))
						}()

					}
				}
			}
		}()
	}
	wg.Wait()
	return message.GetMap()
}
