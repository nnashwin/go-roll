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

func isOp(r rune) bool {
	return r == '+' || r == '-'
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rollDie(dieType int) int {
	return rint.Gen(dieType) + 1
}

func stringContainedInRegex(s string, re *regexp.Regexp) bool {
	if len(re.Find([]byte(s))) == len(s) {
		return true
	}
	return false
}

var opMap = map[string](func(int, int) int){
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
}

func main() {
	re := regexp.MustCompile(`[0-9]*d[0-9]*(\+|\-)*[0-9]+`)

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

		var op string
		for _, ch := range dieStrArr[1] {
			if isOp(ch) {
				op = string(ch)
			}
		}

		var err error
		var numDie int
		//  set numDie to 1 in order to handle case where the number of dice is not specified
		if dieStrArr[0] == "" {
			numDie = 1
		} else {
			numDie, err = strconv.Atoi(dieStrArr[0])
			if err != nil {
				fmt.Printf("The number of die string to int conversion failed with the following error: %s\n", err)
				os.Exit(1)
			}
		}

		// use SplitN because it works when there is an operator and when there is not an operator
		typeStrArr := strings.SplitN(dieStrArr[1], op, len(op)+1)

		typeDie, err := strconv.Atoi(typeStrArr[0])
		if err != nil {
			fmt.Printf("The type of die string to int conversion failed with the following error: %s\n", err)
			os.Exit(1)
		}

		modifierDie := 0
		if len(typeStrArr) > 1 {
			modifierDie, err = strconv.Atoi(typeStrArr[1])
			if err != nil {
				fmt.Printf("The modifier of die string to int conversion failed with the following error: %s\n", err)
				os.Exit(1)
			}
		}

		// handle edge case when a type of die that doesn't exist is thrown
		if numDie == 0 || typeDie == 0 {
			fmt.Println("roll: You ponder the meaning of the number '0'. Consequently, you forget to throw any dice")
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
			// use the op we found in the type of die / modifier string to call the op function on the die
			if opFn, exists := opMap[op]; exists {
				fmt.Printf("Rolling a d%d %s %d:   ", typeDie, op, modifierDie)
				val := maxInt(opFn(rint.Gen(typeDie)+1, modifierDie), 1)
				fmt.Printf("%d\n", val)
			} else {
				// add one in order to offset and mimic a die
				fmt.Printf("Rolling a d%d:    ", typeDie)
				val := rint.Gen(typeDie) + 1
				fmt.Printf("%d\n", val)
			}
		}

		break
	default:
		fmt.Printf("roll: '%s' is not a roll command\nFor help, type: 'roll --help'\n", cmdStr)
		break
	}
}
