package main

import "fmt"

/// Loop examples there is no while loop in Go
// to get infinite loop you can just use for{}
// you can use break to escape loop
// you can use continue to escape next line

func forLoop() int {
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}
	return i
}

func forLoop2() {
	for i := 1; i < 5; i++ {
		fmt.Println("Printing")
	}
}

// basic form when you dont need character code

func range1() {
	str := "test"

	for i := range str {
		fmt.Println(i)
		fmt.Println(str[i]) // returns ascii code
	}
}

// iterates over index where i is index and c is character code
func range2() {
	for i, c := range "test" {
		fmt.Println(i, c)
	}
}

//Blank Identifier if you dont need index
func range3() {
	for _, c := range "test" {
		fmt.Println(c)
	}
}

func main() {

	var i int
	fmt.Println("Please choose the function you want to call ")
	fmt.Scanf("%d", &i)

	switch i {
	case 0:
		forLoop()
	case 1:
		forLoop2()
	case 2:
		range1()
	case 3:
		range2()
	case 4:
		range3()
	default:
		fmt.Println("Unknown function")
	}

}
