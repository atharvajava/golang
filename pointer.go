package main

import "fmt"

var (
	module = 15
	ptr    = &module
)

func main() {
	fmt.Println("This is pointer", ptr, " and the value is ", *ptr)
}
