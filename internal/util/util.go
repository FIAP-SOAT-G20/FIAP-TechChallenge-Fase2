package util

import (
	"os"
	"regexp"
)

func RemoveAllSpaces(s string) string {
	return regexp.MustCompile(`\s`).ReplaceAllString(s, "")
}

func ReadGoldenFile(name string) (string, error) {
	content, err := os.ReadFile("testdata/" + name + ".golden")
	if err != nil {
		return "", err
	}
	return RemoveAllSpaces(string(content)), nil
}
