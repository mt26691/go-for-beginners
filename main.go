package main

import "fmt"

func main() {
	fmt.Println("== Variables, Types & Constants ==")

	variablesDemo()
	zeroValuesDemo()
	basicTypesDemo()
	constantsDemo()
}

// variablesDemo shows the two ways to declare a variable and the
// difference between declaring (:=) and assigning (=).
func variablesDemo() {
	fmt.Println("\n-- Variables --")

	// Short form: := declares and infers the type. Functions only.
	language := "Go"
	fmt.Println("language:", language)

	// Long form with an explicit type, used when you want a type the
	// right-hand side would not infer: 36 alone would be an int.
	var temperature float64 = 36
	fmt.Println("temperature:", temperature)

	// = assigns to an already-declared variable; := would be an error here.
	language = "Golang"
	fmt.Println("language (reassigned):", language)
}

// zeroValuesDemo shows that every variable starts at a sensible default
// when you declare it without a value.
func zeroValuesDemo() {
	fmt.Println("\n-- Zero values --")

	var count int
	var ratio float64
	var name string
	var active bool

	fmt.Printf("int:     %d\n", count)
	fmt.Printf("float64: %g\n", ratio)
	fmt.Printf("string:  %q\n", name)
	fmt.Printf("bool:    %t\n", active)
}

// basicTypesDemo shows Go's common basic types and type inference.
func basicTypesDemo() {
	fmt.Println("\n-- Basic types --")

	count := 42 // int
	var big int64 = 9_000_000_000
	price := 19.99      // float64
	greeting := "héllo" // string (UTF-8)
	ready := true       // bool

	fmt.Printf("%-9v has type %T\n", count, count)
	fmt.Printf("%-9v has type %T\n", big, big)
	fmt.Printf("%-9v has type %T\n", price, price)
	fmt.Printf("%-9q has type %T\n", greeting, greeting)
	fmt.Printf("%-9v has type %T\n", ready, ready)

	// Strings are UTF-8. "héllo" is 6 bytes but 5 runes because é
	// takes two bytes. A byte (uint8) is one raw byte; a rune (int32)
	// is one Unicode code point.
	fmt.Printf("%q: %d bytes, %d runes\n",
		greeting, len(greeting), len([]rune(greeting)))
}

// LogLevel is a small enumeration numbered by iota.
type LogLevel int

const (
	LevelDebug LogLevel = iota // 0
	LevelInfo                  // 1
	LevelWarn                  // 2
	LevelError                 // 3
)

// constantsDemo shows constants and a simple iota enumeration.
func constantsDemo() {
	fmt.Println("\n-- Constants --")

	const pi = 3.14159
	fmt.Println("pi:", pi)

	fmt.Printf("levels: debug=%d info=%d warn=%d error=%d\n",
		LevelDebug, LevelInfo, LevelWarn, LevelError)
}
