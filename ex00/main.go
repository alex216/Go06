package main

import (
	"ft"
	"os"
)

// ////////////////////////////////////////
type boolean int

const (
	yes     boolean = 1
	no      boolean = 0
	EvenMsg         = "I have an even number of arguments"
	OddMsg          = "I have an odd number of arguments"
)

func even(nbr int) int { return 1 - nbr%2 }
func ft_len(args []string) int {
	count := 0
	for range args {
		count++
	}
	return count
}

var lengthOfArg = ft_len(os.Args[1:])

//////////////////////////////////////////

func printStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}

func isEven(nbr int) boolean {
	if even(nbr) == 1 {
		return yes
	} else {
		return no
	}
}

func main() {
	if isEven(lengthOfArg) == 1 {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
