package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// load in sacks
	dat, err := os.ReadFile("sacks.txt")
    if err != nil {
        panic(err)
    }

	// get rucksacks
	sacks := strings.Split(string(dat), "\n")

	// track sum of priorities
	total := 0

	// add up totals
	for i := 1; i < len(sacks); i += 3 {
		items := make(map[string]bool)
	}
	for j, sack := range sacks {

		first := make(map[string]bool)

		for i, item := range sack {

			if (i < len(sack) / 2) {
				first[string(item)] = true
			} else {
				if (first[string(item)]) {
					
					// upper vs lower
					if (item < 91) {
						total += int(item) - 38
					} else {
						total += int(item) - 96
					}
					break
				}
			}
		}
	}

	fmt.Println(total)
}