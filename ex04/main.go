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

func ftMax(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func atoi(s string) int {
	ans := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return -1
		}
		ans = ans*10 + int(c-'0')
	}
	return ans
}

func displayFileAtC(file *os.File, c int64) {
	file_size, err := file.Stat()
	if err != nil {
		ftPutStrLn("ERROR: " + err.Error())
		return
	}
	filesize := file_size.Size()
	buf := make([]byte, c)
	n, err := file.ReadAt(buf, ftMax(0, filesize-c))
	if err != nil && err.Error() != "EOF" {
		ftPutStrLn("ERROR: " + err.Error())
		return
	}
	ftPutStr(string(buf[:n]))
}

func ztail() (res int) {
	c := int64(atoi(os.Args[2]))
	if c == -1 {
		return
	}
	if ftLen(os.Args) == 3 {
		buf := make([]byte, 1024)
		n, err := os.Stdin.Read(buf)
		if err != nil || n == 0 {
			return
		}
		ftPutStr(string(buf[:n]))
	} else {
		filenum := ftLen(os.Args[3:])
		for idx, filename := range os.Args[3:] {
			file, err := os.Open(filename)
			if err != nil {
				ftPutStrLn("ERROR: " + err.Error())
				defer func() {
					res = 1
				}()
				continue
			}
			defer file.Close()
			if filenum > 1 {
				if idx != 0 {
					ft.PrintRune('\n')
				}
				ftPutStr("==> " + filename + " <==\n")
			}
			displayFileAtC(file, c)
		}
	}
	return 0
}

func main() {
	res := ztail()
	if res == 1 {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
