// Package textx holds small string helpers built on the standard library
// and one third-party package.
package textx

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// titler title-cases English text. golang.org/x/text/cases is the
// maintained replacement for the deprecated strings.Title.
var titler = cases.Title(language.English)

// Label turns a raw piece of text into a tidy display label: it trims the
// surrounding spaces and title-cases each word. It calls the unexported
// squeeze helper to collapse runs of inner whitespace first.
func Label(raw string) string {
	cleaned := squeeze(strings.TrimSpace(raw))
	return titler.String(cleaned)
}

// squeeze collapses every run of whitespace into a single space. It is
// unexported, so only code inside package textx can call it.
func squeeze(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
