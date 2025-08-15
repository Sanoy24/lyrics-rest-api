package util

import "strings"

func FormatTsQuery(query string) string {
	splittedString := strings.Fields(query)
	return strings.Join(splittedString, " & ")
}
