package main

import "fmt"

func bestLeagueFinish(finished ...int) int {
	best := finished[0]
	for _, i := range finished {
		if i < best {
			best = i
		}
	}
	return best
}

func main() {
	bestFinish := bestLeagueFinish(13, 10, 12, 15, 16, 18)
	fmt.Println(bestFinish)
}
