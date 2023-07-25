package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func main() {
	if len(os.Args) == 1 {
		return
	}
	// Writing arguments in a single string
	str := os.Args[1]
	if str == "" {
		os.Exit(0)
	}
	for _, v := range os.Args[2:] {
		str += " " + v
		fmt.Println(str)
	}
	//2. Checking weather str contain "\n" or not ---> executing the ascii-art
	prev := 'a'
	severallines := false
	isempty := true
	for _, v := range str {
		if v == 'n' && prev == '\\'{
			severallines=true
		} 
		prev = v
	}
	//3. Writing text line by line into res
	res := ""
	if severallines {
		args := strings.Split(str, "\\n")
		for _, wordii := range args {
			if wordii != "" {
				isempty=false
			} 
		}
		if isempty{
			args = args[1:]
		}
		for _, word := range args {
			if word == "" {
				fmt.Println()
			}else{
			//if word is empty new line 
			for i := 0; i < 8; i++ {
				for _, letter := range word {
						res += GetLine(1 + int(letter-' ')*9+i)						
				}
				fmt.Println(res)
				res = ""
			}
			}
		}

	} else {
		for i := 0; i < 8; i++ {
			for _, letter := range str {
				res += GetLine(1 + int(letter-' ')*9 + i)
			}
			fmt.Println(res)
			res = ""
		}
	}
}

func GetLine(num int) string {
	f, e := os.Open("standard.txt")
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(0)
	}
	defer f.Close()
	content := bufio.NewScanner(f)
	lineNum := 0
	line := ""
	for content.Scan() {
		if lineNum == num {
			line = content.Text()
		}
		lineNum++
	}
	return line
}
