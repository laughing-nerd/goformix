package tests

import (
	"testing"

	"github.com/laughing-nerd/goformix/utils"
)

var testNumber = []struct {
	name     string
	expected string
	val      interface{}
}{
	{"T1", "", 5},
	{"T2", "", 5.5},
	{"T3", "Invalid number", "5"},
	{"T4", "Invalid number", true},
}

func TestValidateNumber(t *testing.T) {
  for _, tc := range testNumber {
    t.Run(tc.name, func(t *testing.T) {
      got := utils.ValidateNumber(tc.val)
      if got != tc.expected {
        t.Errorf("Expected '%s', but got '%s'", tc.expected, got)
      }
    })
  }
}
