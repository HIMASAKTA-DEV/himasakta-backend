package utils

import (
	"fmt"
	"regexp"
	"strings"
)

var tagRegex = regexp.MustCompile(`^#[a-zA-Z0-9-]+$`)

func SanitizeHashtags(input string) (string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", nil
	}

	rawTags := strings.Split(input, ",")
	uniqueTags := make(map[string]bool)
	var finalTags []string

	for _, rawTag := range rawTags {
		tag := strings.TrimSpace(rawTag)
		if tag == "" {
			continue
		}

		if !strings.HasPrefix(tag, "#") {
			return "", fmt.Errorf("invalid hashtag '%s': harus dimulai dengan #. contoh: #news, #event, #its", tag)
		}

		if len(tag) > 20 {
			return "", fmt.Errorf("invalid hashtag '%s': terlalu panjang (max 20 karakter)", tag)
		}

		if !tagRegex.MatchString(tag) {
			return "", fmt.Errorf("invalid hashtag '%s': hanya alphanumeric dan dash (-) yang diizinkan", tag)
		}

		if !uniqueTags[tag] {
			uniqueTags[tag] = true
			finalTags = append(finalTags, tag)
		}
	}

	return strings.Join(finalTags, ","), nil
}
