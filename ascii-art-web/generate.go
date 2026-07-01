package main

import (
	"strings"
)

func Generate(lines []string, banner map[rune][]string) string {
	var b strings.Builder
	for _, line := range lines {
		if line == "" {
			b.WriteString("\n")
			continue
		}
		for i := 0; i < 8; i++ {
			for _, ch := range line {
				if ascii, exist := banner[ch]; exist {
					b.WriteString(ascii[i])
				}
			}
			b.WriteString("\n")
		}
	}
	return b.String()
}
