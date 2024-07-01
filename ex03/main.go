package main

import (
	"ft"
	"os"
)

func ftLen(args []string) int {
	i := 0
	for range args {
		i++
	}
	return i
}

func ftPutStr(str string) {
	for _, c := range str {
		ft.PrintRune(c)
	}
}

func ftPutStrLn(str string) {
	ftPutStr(str)
	ft.PrintRune('\n')
}

func displayFile(filename string) {
	contents, err := os.ReadFile(filename)
	if err != nil {
		ftPutStrLn("ERROR: " + err.Error())
	}
	ftPutStr(string(contents))
}

func main() {
	if ftLen(os.Args) == 1 {
		buf := make([]byte, 42)
		for {
			n, err := os.Stdin.Read(buf)
			if err != nil || n == 0 {
				return
			}
			ftPutStr(string(buf[:n]))
		}
	}
	for _, filename := range os.Args[1:] {
		displayFile(filename)
	}
}
