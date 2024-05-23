package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s, err := os.ReadFile("standard.txt")
	if err != nil {
		os.Exit(1)
	}
	test := os.Args[1]
	txt := strings.Split(string(s[1:]), "\n\n")
	var str [][]string
	for _, v := range txt {
		str = append(str, strings.Split(v, "\n"))
	}
	if len(os.Args) != 2 {
		os.Exit(2)
	}
	for _, v := range test {
		if v < 32 || v > 126 {
			fmt.Println("Invalid input")
			os.Exit(1)
		}
	}
	printTxt(test, str)
	if strings.HasSuffix(test, "\\n") {
		if len(test) != 2 {
			fmt.Println()
		}
	}
}

func printTxt(s string, sl [][]string) {
	var str [][]string
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && s[i] == '\\' && s[i+1] == 'n' {
			for j := 0; j < 8; j++ {
				for k := 0; k < len(str); k++ {
					fmt.Print(str[k][j])
				}
				fmt.Println()
				if len(str) == 0 {
					break
				}
			}
			i++
			str = [][]string{}
		} else {
			str = append(str, sl[rune(s[i])-32])
		}
	}
	if len(str) > 0 {
		for j := 0; j < 8; j++ {
			for k := 0; k < len(str); k++ {
				fmt.Print(str[k][j])
			}
			fmt.Println()
		}
	}
}
