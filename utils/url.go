package utils

import (
	"regexp"
	"strings"
)

func ValidateUrl(allowedDomains []string, val string) string {
	pattern := `^(https|http)?://`

	if len(allowedDomains) != 0 {
		domains := strings.Join(allowedDomains, "|")
		pattern += `(` + domains + `)$`
	} else {
		pattern += `[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	}

	re := regexp.MustCompile(pattern)
	if !re.MatchString(val) {
		return "Invalid URL"
	}
	return ""
}
