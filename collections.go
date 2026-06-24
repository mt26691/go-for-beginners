package main

import (
	"fmt"
	"sort"
)

// slicesDemo shows how slices grow with append, how len and cap differ, a slice
// expression, ranging over a slice, and the aliasing gotcha that happens when
// two slices share the same backing array.
func slicesDemo() {
	fmt.Println("\n-- Slices --")

	// make a slice with length 0 but room (capacity) for 2 Tasks, then append.
	tasks := make([]Task, 0, 2)
	tasks = append(tasks, Task{ID: 1, Title: "write code"})
	tasks = append(tasks, Task{ID: 2, Title: "run tests"})
	tasks = append(tasks, Task{ID: 3, Title: "ship it"})

	fmt.Printf("len=%d cap=%d after three appends\n", len(tasks), cap(tasks))

	for i, t := range tasks {
		fmt.Printf("  tasks[%d] = %s\n", i, t.Summary())
	}

	// a slice expression: middle holds tasks[1] and tasks[2], sharing the same
	// backing array as tasks.
	middle := tasks[1:3]
	fmt.Printf("middle := tasks[1:3] -> len=%d cap=%d\n", len(middle), cap(middle))

	// aliasing: middle and tasks point at the same array, so writing through one
	// is visible through the other.
	middle[0].Title = "RUN TESTS (edited via middle)"
	fmt.Println("after middle[0].Title = ...:")
	fmt.Println("  tasks[1] =", tasks[1].Summary())
	fmt.Println("  middle[0] =", middle[0].Summary())
}

// mapsDemo shows a map[int]Task keyed by ID: set and get values, the comma-ok
// lookup for present and absent keys, delete, and why we iterate by sorted keys
// (map iteration order is unspecified, so raw range order is not reproducible).
func mapsDemo() {
	fmt.Println("\n-- Maps --")

	// a map literal keyed by task ID. This is the shape of the in-memory store
	// you build later in the course: map[int]Task.
	byID := map[int]Task{
		1: {ID: 1, Title: "write code"},
		2: {ID: 2, Title: "run tests"},
		3: {ID: 3, Title: "ship it", Done: true},
	}
	byID[4] = Task{ID: 4, Title: "deploy"}

	// comma-ok: ok tells you whether the key was present, separate from the value.
	if t, ok := byID[2]; ok {
		fmt.Println("found id 2:", t.Summary())
	}
	if _, ok := byID[99]; !ok {
		fmt.Println("id 99 is not in the map (comma-ok said so)")
	}

	delete(byID, 4)
	fmt.Printf("after delete(byID, 4): %d tasks remain\n", len(byID))

	// map iteration order is random, so sort the keys for stable output.
	keys := make([]int, 0, len(byID))
	for id := range byID {
		keys = append(keys, id)
	}
	sort.Ints(keys)

	fmt.Println("tasks in sorted-key order:")
	for _, id := range keys {
		fmt.Printf("  %d -> %s\n", id, byID[id].Summary())
	}
}

// nilMapDemo shows that reading from a nil map returns the zero value safely,
// while writing to a nil map panics.
func nilMapDemo() {
	fmt.Println("\n-- Nil maps --")

	var tasks map[int]Task // nil map: no underlying storage yet

	// reading a missing key from a nil map is safe: you get the zero value.
	t, ok := tasks[1]
	fmt.Printf("read from nil map: ok=%t, value=%+v\n", ok, t)

	// writing would panic, so we do not run it:
	//   tasks[1] = Task{ID: 1} // panic: assignment to entry in nil map
	fmt.Println("writing to a nil map would panic; make() the map first")
}
