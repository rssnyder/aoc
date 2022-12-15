package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {

	// load in snack list
	dat, err := os.ReadFile("rations.txt")
    if err != nil {
        panic(err)
    }

	// get list of entries
	elfs := strings.Split(string(dat), "\n\n")

	// track elves and largest
	most := 0

	// add up totals
	for _, elf := range elfs {
		elfTotal := 0

		// add up all snacks per elf
		for _, snack := range strings.Split(elf, "\n") {
			value, err := strconv.Atoi(snack)
			if err != nil {
				panic(err)
			}
			elfTotal += value
		}

		// check if elf has the most
		if (elfTotal > most) {
			most = elfTotal
		}
	}

	fmt.Println(most)
}