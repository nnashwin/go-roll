package main

import (
	"fmt"
	rint "github.com/tlboright/go-rint"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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

		if !stringContainedInRegex(cmdStr, re) {
			fmt.Printf("roll: specified an invalid dice expression\nFor help, type: 'roll --help'\n")
			return
		}

		rint.Init()
		dieStrArr := strings.Split(cmdStr, "d")

		var rollVals []int

		numDie, err := strconv.Atoi(dieStrArr[0])
		if err != nil {
			fmt.Errorf("The number of die string to int conversion failed with the following error: %s", err)
		}

		typeDie, err := strconv.Atoi(dieStrArr[1])
		if err != nil {
			fmt.Errorf("The type of die string to int conversion failed with the following error: %s", err)
		}

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
