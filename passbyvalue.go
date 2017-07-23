package main

import "fmt"

func main() {
	name := "Atharva"
	course := "Learning golang"

	fmt.Println("\nHi", name, "you're currently watching", course)
	changeCourse(course)
	fmt.Println("\nHi", name, "you're currently watching", course)
}

func changeCourse(course string) string {
	course = "First Look: Native Docker clustering"
	fmt.Println("\nTrying to change course to", course)
	return course
}
