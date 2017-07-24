/*
Why Structs?
-- Structs at a high levl. structs let uis define custom dataType.
eg :
type Circle struct {
	r float64
	d float64
	c float64
}

Object Oriented Programming in GO ...

It is a unix philosiphy of creating small modules and using it for only one purpose
Go doesn't have object type
neither it has class
go also doesn't support Inheritance

Go is not designed for object oriented programming
*/

package main

import(
	"fmt"
)

func main(){
	type courseMeta struct {
		Author  string
		Level string
		Rating float64
	}

	//var DockerDeepDive courseMeta
	//DockerDeepDive:= new(courseMeta) // using new keyword gives us a pointer 
	// The above two method initialize struct with 0

	DockerDeepDive := courseMeta{
		Author : "Atharva",
		Level : "Intermediate",
		Rating: 5,
	}

	fmt.Println(DockerDeepDive.Author)	
	DockerDeepDive.Rating=1
	fmt.Println(DockerDeepDive.Rating)

}