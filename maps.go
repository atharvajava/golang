package main

import (
	"fmt"
)

func main(){
	// Basic Syntax map[keyType] valueType keytype has to be anything 
	//that can be compared i.e we cant have arrays or slices in keytype
	leagueTitles:= make(map[string]int)
	leagueTitles["Sunderland"] = 6
	leagueTitles["Newcastle"] = 4

	//composite litrals
	recentHead2Head := map[string] int {
		"Sunderland":5,
		"Newcastle":0,
	}	

	fmt.Printf("\n League titles : %v \n Recent head to heads : %v \n", leagueTitles, recentHead2Head)
	iteratingMaps()
	crudMap()
}

//We must treat maps as unordered lists
//It forces programmers to good coding practices
// SImilar to slices Maps are refernce types
// Any changes made to maps are visible to caller
// Maps are not safe for concurrency
// passing Maps is cheap and fast
// as under the hood it is just pointer
// specify size for large maps
// eg : Make(map[<keyType>]<valueType>, size) it may incrase performance

func iteratingMaps(){
	testMap := map[string]int {"A":1, "B":2, "C":3, "D":4 , "E":5}

	for key, value:= range testMap{
		fmt.Println("Key is :",key,"Value is :",value)
	}
}

func crudMap(){
	testMap := map[string]int {"A":1, "B":2, "C":3, "D":4 , "E":5}
	testMap["A"] = 100
	testMap["F"] = 1992
	fmt.Println(testMap)

	delete(testMap, "F")
	fmt.Println(testMap)
}

//maps are go implementation of hashtables