package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var arguments = os.Args[1:]

	var data *os.File
	var error error

	for i := 0; i < len(arguments); i++ {
		if arguments[i] == "-f" {
			data, error = os.Open(arguments[i+1])
			if error != nil {
				fmt.Println("ayak: failed to open the file!")
				return
			}
		}
	}

	file_scanner := bufio.NewScanner(data)
	file_scanner.Split(bufio.ScanLines)
	var lines []string
	for file_scanner.Scan() {
		lines = append(lines, file_scanner.Text())
	}
	data.Close()

	succes := lex(lines)
	inter(succes)
}
