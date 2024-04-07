package utils

import (
	"regexp"
	"strings"
)

func ValidateEmail(allowedDomains []string, val string) string {
	pattern := `^[a-zA-Z0-9._%+-]+@`

	if len(allowedDomains) != 0 {
		domains := strings.Join(allowedDomains, "|")
		pattern += `(` + domains + `)$`
	} else {
		pattern += `[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	}

	re := regexp.MustCompile(pattern)
	if !re.MatchString(val) {
		return "Invalid email address"
	}
  return ""
}
