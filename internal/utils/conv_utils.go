package utils

import (
	"strings"
)

func ToSlug(s string) string {
	s = strings.ToLower(s)
	var b strings.Builder

	spaceFound := false
	for _, r := range s {
		switch {
		case (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9'):
			if spaceFound {
				b.WriteRune('-')
				spaceFound = false
			}
			b.WriteRune(r)
		case r == ' ':
			spaceFound = true
		}
	}

	return b.String()
}

func ExtractAcronym(s string) string {
	start := strings.Index(s, "(")
	end := strings.Index(s, ")")

	if start == -1 || end == -1 || start > end {
		return ""
	}

	return s[start+1 : end]
}
