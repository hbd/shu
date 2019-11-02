// Package md provides an API for parsing markdown.

package mp

import (
	"strings"
)

// Markdown is a markdown document.
type Markdown struct {
	LinesRaw []string
	Lines    []Line
}

type Line struct {
	Indented bool
	Raw      string
	Str      string
}

// Indented returns true if the given text begins with an indentation.
// Refer to the spec for the definition of an indentation.
func Indented(md string) bool {
	if len(md) < 1 {
		return false
	}
	if string(md[0]) == "\t" {
		return true
	}
	if string(md[0]) == " " {
		var count int
		for _, c := range md {
			if string(c) == " " {
				count++
			} else if string(c) == "\t" {
				return true
			} else {
				// Cannot be indented if not a space or tab.
				return false
			}
			if count == 4 {
				return true
			}
		}
	}

	return false
}

// Parse parses.
func Parse(md string) Markdown {
	ret := Markdown{}

	lines := strings.Split(md, "\n")
	ret.LinesRaw = lines

	for _, line := range lines {
		if len(line) < 1 {
			continue
		}

		l := Line{
			Indented: Indented(line),
			Raw:      line,
			Str:      strings.TrimLeft(line, " \t"),
		}
		ret.Lines = append(ret.Lines, l)
	}
	return ret
}

/* Markdown Spec

# x1 -> x6

*/
