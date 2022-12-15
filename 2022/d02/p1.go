package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// scoring                 Rock   Paper   Scissors
	opponent := map[string]int{"A": 1, "B": 2, "C": 3}
	response := map[string]int{"X": 1, "Y": 2, "Z": 3}

	// load in strategy
	dat, err := os.ReadFile("guide.txt")
    if err != nil {
        panic(err)
    }

	// get rounds
	rounds := strings.Split(string(dat), "\n")

	// track score
	score := 0

	// add up totals
	for _, round := range rounds {
		play := strings.Split(round, " ")
		
		switch response[play[1]] - opponent[play[0]] {
		case 0:
			score += 3 + response[play[1]]
		case 1:
			score += 6 + response[play[1]]
		case -2:
			score += 6 + response[play[1]]
		case -1:
			score += response[play[1]]
		case 2:
			score += response[play[1]]
		}
	}

	fmt.Println(score)
}