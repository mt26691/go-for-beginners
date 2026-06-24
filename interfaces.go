package main

import (
	"fmt"
	"os"
)

// interfacesDemo runs the same code against two different stores through one
// interface, then shows type assertions, a type switch, and the empty interface.
func interfacesDemo() {
	fmt.Println("\n-- Interfaces --")

	// runStore depends on the TaskStore interface, so it works with either store.
	fmt.Println("InMemoryStore:")
	mem := NewInMemoryStore()
	runStore(mem)

	// Wrap the same kind of store in a LoggingStore and run the very same code.
	// The [log] lines come from the wrapper; the results are identical.
	fmt.Println("LoggingStore (wraps an InMemoryStore):")
	logged := NewLoggingStore(NewInMemoryStore(), os.Stdout)
	runStore(logged)

	// A type assertion pulls the concrete type back out of an interface value.
	// The comma-ok form never panics: ok tells you whether the assertion held.
	var store TaskStore = mem
	if im, ok := store.(*InMemoryStore); ok {
		fmt.Printf("type assertion: it really is an *InMemoryStore with %d tasks\n", len(im.All()))
	}

	// The empty interface, any, holds a value of any type. describe uses a type
	// switch to recover the real type at runtime. Reach for any sparingly.
	fmt.Println("describe(42):        ", describe(42))
	fmt.Println("describe(\"hello\"):    ", describe("hello"))
	fmt.Println("describe(a Task):     ", describe(Task{ID: 7, Title: "demo"}))
}

// describe inspects a value of the empty interface type (any) at runtime using a
// type switch. any means "any type at all"; reach for it sparingly, because you
// lose compile-time type safety and have to assert the real type back out.
func describe(v any) string {
	switch x := v.(type) {
	case int:
		return fmt.Sprintf("an int (%d)", x)
	case string:
		return fmt.Sprintf("a string (%q)", x)
	case Task:
		return "a Task: " + x.String()
	default:
		return fmt.Sprintf("some other type: %T", x)
	}
}
