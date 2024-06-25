package main

import (
	"ex02/ft"
	"os"
)

func ft_len(args []string) int {
	i := 0
	for range args {
		i++
	}
	return i
}

func ft_putstr(str string) {
	for _, c := range str {
		ft.PrintRune(c)
	}
}

func ft_putstrnl(str string) {
	ft_putstr(str)
	ft.PrintRune('\n')
}

// unicode not supported if use os.Open then os.Read since it's difficult even tenious to check
// whether byte array contains ascii or unicode without using unicode package
func main() {
	if ft_len(os.Args) != 2 {
		ft_putstrnl("too many arguments")
		return
	}
	contents, err := os.ReadFile(os.Args[1])
	if err != nil {
		ft_putstrnl("file open error")
		return
	}
	ft_putstr(string(contents))
}
