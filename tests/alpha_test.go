package tests

import (
	"testing"

	"github.com/laughing-nerd/goformix/utils"
)

var testAlpha = []struct {
	name     string
	expected string
	args     []string
	val      string
}{
	{"T1", "", []string{}, "Hello World!"},
	{"T2", "", []string{"6", "9"}, "OrangeCar"},
	{"T3", "String length is not within the specified range", []string{"6", "9"}, "Hello"},
	{"T4", "String length is not within the specified range", []string{"6", "9"}, "Hello World!, I am learning Go"},
	{"T5", "", []string{"10"}, "Hello World!, I am learning Go"},
	{"T6", "String length is less than the minimum length", []string{"10"}, "Hi"},
	{"T7", "String length is less than the minimum length", []string{"10", "5"}, "Hi"},
}

func TestValidateString(t *testing.T) {
	for _, tc := range testAlpha {
		t.Run(tc.name, func(t *testing.T) {
			got := utils.ValidateString(tc.args, tc.val)
			if got != tc.expected {
				t.Errorf("Expected '%s', but got '%s'", tc.expected, got)
			}
		})
	}
}
