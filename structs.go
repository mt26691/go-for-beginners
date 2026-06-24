package main

import "fmt"

// structsDemo shows three ways to create a Task and the value-receiver vs
// pointer-receiver mutation contrast.
func structsDemo() {
	fmt.Println("\n-- Structs --")

	literal := Task{ID: 1, Title: "write code", Done: false}
	pointer := &Task{ID: 2, Title: "run tests"}
	fresh := new(Task)
	fresh.ID = 3
	fresh.Title = "ship it"

	fmt.Println(literal.Summary())
	fmt.Println(pointer.Summary())
	fmt.Println(fresh.Summary())

	fmt.Println("\n-- Value vs Pointer Receiver --")

	insideCopy := literal.markDoneByValue()
	fmt.Printf("markDoneByValue set the copy to %t, but...\n", insideCopy)
	fmt.Println("after markDoneByValue:", literal.Summary())

	literal.MarkDone()
	fmt.Println("after MarkDone:       ", literal.Summary())
}

// pointersDemo shows what a pointer is: the address of a value. & takes an
// address, * reads through it, and a nil pointer points at nothing.
func pointersDemo() {
	fmt.Println("\n-- Pointers --")

	count := 41
	p := &count
	*p = *p + 1
	fmt.Printf("count is now %d (changed through the pointer)\n", count)

	var missing *Task
	fmt.Println("a nil *Task prints as:", missing)
}

// labeledTask embeds Task. Embedding is Go's composition: a labeledTask has all
// of Task's fields and methods promoted to it, plus its own Label field.
type labeledTask struct {
	Task
	Label string
}

// embeddingDemo builds a labeledTask and calls the promoted Summary method and
// accesses the promoted Title field directly.
func embeddingDemo() {
	fmt.Println("\n-- Embedding --")

	lt := labeledTask{
		Task:  Task{ID: 4, Title: "review PR", Done: true},
		Label: "urgent",
	}

	fmt.Printf("%s (label: %s)\n", lt.Summary(), lt.Label)
	fmt.Println("promoted field Title:", lt.Title)
}
