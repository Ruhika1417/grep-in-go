package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	in := bufio.NewScanner(file)

	var pattern string
	fmt.Print("enter the string pattern to search: ")
	fmt.Scanln(&pattern)

	if args := os.Args[1:]; len(args) == 1 {
		pattern = args[0]
	}

	for in.Scan() {
		s := in.Text()
		if strings.Contains(s, pattern) {
			fmt.Println(s)
		}
	}
}
