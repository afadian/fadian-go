package util

import "strings"

func ReplaceString(s string, replace map[string]string) string {
	for k, v := range replace {
		s = strings.ReplaceAll(s, k, v)
	}

	return s
}
