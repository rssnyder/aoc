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

	// track largest
	ledger := [3]int{0, 0, 0}

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
		pushOut(&ledger, elfTotal)
	}

	fmt.Println(ledger[0] + ledger[1] + ledger[2])
}

func pushOut(largest *[3]int, newbie int) {
	switch {
	case newbie > largest[0]:
		largest[2] = largest[1]
		largest[1] = largest[0]
		largest[0] = newbie
	case newbie > largest[1]:
		largest[2] = largest[1]
		largest[1] = newbie
	case newbie > largest[2]:
		largest[2] = newbie
	}
}