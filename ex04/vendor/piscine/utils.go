package piscine

import (
	"ft"
)

func FtLen(args []string) int {
	i := 0
	for range args {
		i++
	}
	return i
}

func FtPutStr(str string) {
	for _, c := range str {
		ft.PrintRune(c)
	}
}

func FtPutStrLn(str string) {
	FtPutStr(str)
	ft.PrintRune('\n')
}

func FtMax(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func Atoi(s string) int {
	ans := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return -1
		}
		ans = ans*10 + int(c-'0')
	}
	return ans
}

func Strcmp(s1 string, s2 string) int {
	if s1 == s2 {
		return 0
	} else if s1 > s2 {
		return 1
	} else {
		return -1
	}
}
