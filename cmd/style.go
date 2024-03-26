package cmd

import (
	"log"
	"strings"
	"unicode"
)

func getStyleOrDefault(enumStyle string, defaultStyle enumStyle) enumStyle {
	if enumStyle == "" {
		return defaultStyle
	}
	for _, e := range allEnumStyles {
		if string(e) == enumStyle {
			return e
		}
	}
	log.Fatalf("invalid enum style %s! options: %s", enumStyle, strings.Join(getValidEnumsStylesAsStrings(), ","))
	return defaultStyle
}

func mapStringSlices(stringSlices []string, mapFunc func(toMap string) string) []string {
	mapped := make([]string, len(stringSlices))
	for i, s := range stringSlices {
		mapped[i] = mapFunc(s)
	}
	return mapped
}

func splitValueUsingCase(value string) []string {
	var valueChunks []string
	runes := []rune(value)
	currentChunk := ""
	for _, r := range runes {
		if unicode.IsUpper(r) {
			valueChunks = append(valueChunks, currentChunk)
			currentChunk = string(unicode.ToLower(r))
		} else {
			currentChunk += string(r)
		}
	}
	if len(currentChunk) > 0 {
		valueChunks = append(valueChunks, currentChunk)
	}
	return valueChunks
}

func applyEnumStyle(valuesSlice []string, style enumStyle) []string {
	if style == enumStyleCamelCase {
		return valuesSlice
	}
	mappingFunc := func(toMap string) string { return toMap }
	switch style {
	case enumStyleScreamingSnakeCase, enumStyleScreamingKebabCase, enumStyleUpperCase:
		mappingFunc = strings.ToUpper
	case enumStylePascalCase:
		mappingFunc = strings.Title
	}

	separator := ""
	switch style {
	case enumStyleSnakeCase, enumStyleScreamingSnakeCase:
		separator = "_"
	case enumStyleKebabCase, enumStyleScreamingKebabCase:
		separator = "-"
	}

	toReturn := make([]string, len(valuesSlice))
	for i, v := range valuesSlice {
		upperChunks := mapStringSlices(splitValueUsingCase(v), mappingFunc)
		toReturn[i] = strings.Join(upperChunks, separator)
	}
	return toReturn
}
