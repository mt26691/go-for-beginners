package main

import (
	"fmt"
	"os"
)

// interfacesDemo runs the same code against two different stores through one
// interface, then shows type assertions, a type switch, and the empty interface.
func interfacesDemo() {
	fmt.Println("\n-- Interfaces --")

	// TODO: build an *InMemoryStore and run it through runStore.
	mem := NewInMemoryStore()
	runStore(mem)

	// TODO: wrap the same store in a *LoggingStore and run that too — same code,
	// a different TaskStore underneath.
	logged := NewLoggingStore(NewInMemoryStore(), os.Stdout)
	runStore(logged)

	// TODO: show type assertions, a type switch, and the empty interface (any).
	fmt.Println(describe(42))
}

// describe inspects a value of the empty interface type (any) at runtime using a
// type switch. any means "any type at all"; reach for it sparingly, because you
// lose compile-time type safety and have to assert the real type back out.
func describe(v any) string {
	// TODO: switch on v.(type) and return a short description per case.
	return fmt.Sprintf("describe: %v", v)
}
