package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// scoring                 Rock   Paper   Scissors
	response := map[string]int{"X": 1, "Y": 2, "Z": 3}
	winner := map[string]string{"A": "Y", "B": "Z", "C": "X"}
	loser := map[string]string{"A": "Z", "B": "X", "C": "Y"}
	draw := map[string]string{"A": "X", "B": "Y", "C": "Z"}


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
		
		switch play[1] {
		case "Y":
			fmt.Println("draw")
			score += 3 + response[draw[play[0]]]
		case "Z":
			fmt.Println("win")
			score += 6 + response[winner[play[0]]]
		case "X":
			fmt.Println("loss")
			score += response[loser[play[0]]]
		}
	}

	fmt.Println(score)
}