package main

import "fmt"

// structsDemo will show three ways to create a Task and the value-receiver vs
// pointer-receiver mutation contrast.
func structsDemo() {
	fmt.Println("\n-- Structs --")

	// TODO: also create a Task with &Task{} and with new, and print each Summary()
	t := Task{ID: 1, Title: "write code"}
	fmt.Println(t.Summary())

	// TODO: print the result before and after each call to show the contrast
	_ = t.markDoneByValue()
	t.MarkDone()
}

// pointersDemo will show what a pointer is: the address of a value, with & and
// *, plus a nil pointer.
func pointersDemo() {
	fmt.Println("\n-- Pointers --")
	// TODO: take the address of an int with &, change it through *, print it
	// TODO: declare a nil *Task and print it
}

// embeddingDemo will build a struct that embeds Task and call its promoted
// method and field.
func embeddingDemo() {
	fmt.Println("\n-- Embedding --")
	// TODO: define a struct that embeds Task, build one, call the promoted
	// Summary() and read the promoted Title field
}
