package golib

import (
    "strings"
)

func GetFinalArgs(slice []string, start int) string {
	result := strings.Join(slice[start:len(slice)], " ")
	return result
}
