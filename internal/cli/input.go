package cli

import "strings"

func CleanInput(text string) []string {
	cleanString := strings.Trim(text, " ")
	cleanString = strings.ToLower(cleanString)
	for strings.Contains(cleanString, "  ") {
		cleanString = strings.ReplaceAll(cleanString, "  ", " ")
	}
	splitString := strings.Split(cleanString, " ")
	return splitString
}
