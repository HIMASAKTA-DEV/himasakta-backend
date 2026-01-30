package mylog

import (
	"bytes"
	"fmt"
	"strings"
)

func ColorizeINFOJSON(jsonStr string) string {
	var prettyJSON bytes.Buffer
	for _, line := range strings.Split(jsonStr, "\n") {
		prettyJSON.WriteString(ColorizeInfo(line) + "\n")
	}
	return prettyJSON.String()
}

func ColorizeErrorJSON(jsonStr string) string {
	var prettyJSON bytes.Buffer
	for _, line := range strings.Split(jsonStr, "\n") {
		prettyJSON.WriteString(ColorizeError(line) + "\n")
	}
	return prettyJSON.String()
}

func ColorizePanicJSON(jsonStr string) string {
	var prettyJSON bytes.Buffer
	for _, line := range strings.Split(jsonStr, "\n") {
		prettyJSON.WriteString(ColorizePanic(line))
	}
	return prettyJSON.String()
}

func ColorizeInfo(s string) string {
	return fmt.Sprintf("\033[36m%s\033[0m", s)
}

func ColorizeError(s string) string {
	return fmt.Sprintf("\033[31m%s\033[0m", s)
}

func ColorizePanic(s string) string {
	return fmt.Sprintf("\033[33m%s\033[0m\n", s)
}
