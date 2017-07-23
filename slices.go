package main

import "fmt"

func main() {
	//Declares a slice called myCourses
	myCourses := make([]string, 5, 10)

	fmt.Println("Length is :", len(myCourses), "\n Capacity is :", cap(myCourses))
}
