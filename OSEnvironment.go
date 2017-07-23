package main

import "fmt"
import "os"

func main() {
	name := os.Getenv("USER")
	in, config := os.Open("selinuxConfig")

	for _, env := range os.Environ() {
		fmt.Println(env)
	}
	fmt.Println("Hello ", name)
	fmt.Println("Config", in, config)
}
