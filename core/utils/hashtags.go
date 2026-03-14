package utils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	"github.com/google/uuid"
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

func SplitHashTags(hashtags string) ([]entity.Tag, error) {
	rawTags := strings.Split(hashtags, ",")

	var tags []string

	for _, t := range rawTags {
		t = strings.TrimSpace(t)       // hilangkan spasi
		t = strings.TrimPrefix(t, "#") // hilangkan # di depan
		if t != "" {
			tags = append(tags, t)
		}
	}

	var tagEntities []entity.Tag

	for _, t := range tags {
		tagEntities = append(tagEntities, entity.Tag{
			Id:   uuid.New(),
			Name: t,
		})
	}

	return tagEntities, nil
}
