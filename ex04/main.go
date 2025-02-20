package main

import (
	"ft"
	"os"
	"piscine"
)

func validateArgs() (int64, int) {
	if piscine.FtLen(os.Args) < 4 {
		return -1, 1
	}
	if piscine.Strcmp(os.Args[1], "-c") != 0 {
		return -1, 1
	}
	c := int64(piscine.Atoi(os.Args[2]))
	if c == -1 {
		return -1, 1
	}
	return c, 0
}

func displayFileAtC(file *os.File, c int64) {

	file_size, err := file.Stat()
	if err != nil {
		piscine.FtPutStrLn(err.Error())
		return
	}

	filesize := file_size.Size()
	buf := make([]byte, c)
	n, err := file.ReadAt(buf, piscine.FtMax(0, filesize-c))
	if err != nil && err.Error() != "EOF" {
		piscine.FtPutStrLn(err.Error())
		return
	}
	piscine.FtPutStr(string(buf[:n]))
}

func ztail() (res int) {

	c, err := validateArgs()
	if err == 1 {
		return
	}

	filenum := piscine.FtLen(os.Args[3:])
	for i, filename := range os.Args[3:] {
		validfilenum := 0
		file, err := os.Open(filename)
		if err != nil {
			piscine.FtPutStrLn(err.Error())
			defer func() {
				res = 1
			}()
		} else {
			defer file.Close()
			if filenum > 1 {
				if i > 0 || validfilenum == 1 {
					ft.PrintRune('\n')
				}
				piscine.FtPutStr("==> " + filename + " <==\n")
			}
			displayFileAtC(file, c)
			validfilenum = 1
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
