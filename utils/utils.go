// Package utils provides util functions.
package utils

import (
	"bufio"
	"strings"
)

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func NewScanner(input string) *bufio.Scanner {
	return bufio.NewScanner(strings.NewReader(input))
}
