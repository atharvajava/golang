package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 5
	y := 2
	str := "Hello World " + fmt.Sprint(x)
	str2 := "Hello World " + strconv.Itoa(y)
	fmt.Println(str)
	fmt.Println(str2)
}
