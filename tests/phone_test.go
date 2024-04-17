package tests

import (
	"testing"

	"github.com/laughing-nerd/goformix/utils"
)

var testPhone = []struct {
	name     string
	expected string
	args     []string
	val      string
}{
	{"T1", "", []string{}, "9876543210"},
}

func TestValidatePhone(t *testing.T) {
	for _, tc := range testPhone {
		out := utils.ValidatePhone(tc.args, tc.val)
		if out != tc.expected {
			t.Errorf("Expected '%s' but got '%s'", tc.expected, out)
		}
	}
}
