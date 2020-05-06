package main

import (
	"fmt"
	rint "github.com/tlboright/go-rint"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const DIE_SIZE_LIMIT = 10000

func rollDie(dieType int) int {
	return rint.Gen(dieType) + 1
}

func stringContainedInRegex(s string, re *regexp.Regexp) bool {
	if len(re.Find([]byte(s))) == len(s) {
		return true
	}
	return false
}

func main() {
	re := regexp.MustCompile(`^[0-9]+d[0-9]+`)

	if len(os.Args) < 2 {
		fmt.Printf("roll: A dice roll wasn't made\nFor help, type: 'roll --help'\n")
		os.Exit(1)
	}

	switch cmdStr := os.Args[1]; {
	case len(re.Find([]byte(cmdStr))) > 0:

		// if the string is not contained entirely in the regex, it must be an invalid roll string
		// ensures we don't take any unexpected behaviors for half matches
		if !stringContainedInRegex(cmdStr, re) {
			fmt.Printf("roll: specified an invalid dice expression\nFor help, type: 'roll --help'\n")
			return
		}

		// initialize the prng only after ensuring that the imput is correct
		dieStrArr := strings.Split(cmdStr, "d")

		var rollVals []int

		numDie, err := strconv.Atoi(dieStrArr[0])
		if err != nil {
			fmt.Printf("The number of die string to int conversion failed with the following error: %s", err)
			os.Exit(1)
		}

		typeDie, err := strconv.Atoi(dieStrArr[1])
		if err != nil {
			fmt.Printf("The type of die string to int conversion failed with the following error: %s", err)
			os.Exit(1)
		}

		if numDie > DIE_SIZE_LIMIT {
			fmt.Println("roll: You can't fit this many dice in your hands.  Shorten your query and try again")
			os.Exit(1)
		}

		if typeDie > DIE_SIZE_LIMIT {
			fmt.Println("roll: You can't conceive of a die this large.  Shorten your query and try again")
			os.Exit(1)
		}

		rint.Init()

		for i := 0; i < numDie; i++ {
			// add one in order to offset and mimic a die
			rollVals = append(rollVals, rint.Gen(typeDie)+1)
		}

		fmt.Println(rollVals)
		break
	default:
		fmt.Printf("roll: '%s' is not a roll command\nFor help, type: 'roll --help'\n", cmdStr)
		break
	}
}
