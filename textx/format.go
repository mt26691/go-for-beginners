// Package textx holds small string helpers built on the standard library.
package textx

// Label cleans up a raw piece of text into a tidy display label.
// On the finish branch it trims spaces and title-cases the words.
func Label(raw string) string {
	// TODO: trim the surrounding spaces (strings.TrimSpace) and
	// title-case the words. The actual title-casing will come from a
	// third-party package you add with `go get`.
	return raw
}
