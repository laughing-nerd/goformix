package tests

import (
	"testing"

	"github.com/laughing-nerd/goformix/utils"
)

var testUrl = []struct {
	name     string
	expected string
	args     []string
	val      string
}{
	{"T1", "", []string{}, "https://www.github.com"},
	{"T2", "Invalid URL", []string{}, "www.github.com"},
	{"T3", "Invalid URL", []string{"google.com"}, "https://github.com"},
	{"T4", "", []string{"github.com"}, "https://github.com"},
	{"T5", "Invalid URL", []string{"github"}, "https://github.com"},
}

func TestValidateEmail(t *testing.T) {
	for _, tc := range testUrl {
		out := utils.ValidateUrl(tc.args, tc.val)
		if out != tc.expected {
			t.Errorf("Expected '%s' but got '%s'", tc.expected, out)
		}
	}
}
