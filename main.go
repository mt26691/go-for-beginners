package main

import "fmt"

func main() {
	fmt.Println("== Structs, Methods & Pointers ==")

	structsDemo()
	pointersDemo()
	embeddingDemo()

	fmt.Println("\n== Slices & Maps ==")

	slicesDemo()
	mapsDemo()
	nilMapDemo()

	fmt.Println("\n== Interfaces ==")

	interfacesDemo()

	fmt.Println("\n== Errors ==")

	errorsDemo()

	fmt.Println("\n== Concurrency ==")

	concurrencyDemo()
}
