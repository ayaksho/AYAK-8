package main

import (
	"fmt"
	"strconv"
	"strings"
)

func inter(lexed lexer_out) {
	variables := make(map[string]string)

	functions := lexed.functions
	words := lexed.words

	for i := 0; i < len(words); i++ {
		arguments := words[i].function_arguments
		switch words[i].function_type {
		case _mov:
			if arguments[0] == ax || arguments[0] == bx || arguments[0] == cx {
				if arguments[1][0] == '$' {
					name := strings.Replace(arguments[1], "$", "", 1)
					variables[arguments[0]] = strings.ReplaceAll(variables[name], string('"'), "")
				} else {
					variables[arguments[0]] = arguments[1]
				}
			}
		case _int:
			switch arguments[0] {
			case io:
				if variables[bx] == "0" {
					dl := "\\" + "n"
					fmt.Print(strings.ReplaceAll(variables[ax], dl, "\n"))
				}
			}
		case _jmp:
			i = functions[arguments[0]]
		case _db:
			variables[arguments[0]] = arguments[1]
		case _add:
			integer, error := strconv.Atoi(variables[strings.Replace(arguments[0], "$", "", 1)])
			if error != nil {
				fmt.Println("ayak: failed to convert string to integer!")
				return
			}
			variables[strings.Replace(arguments[0], "$", "", 1)] = strconv.Itoa(integer + 1)
		case _addx:
			integer, error1 := strconv.Atoi(variables[strings.Replace(arguments[0], "$", "", 1)])
			add, error2 := strconv.Atoi(arguments[1])
			if arguments[1][0] == '$' {
				add, error2 = strconv.Atoi(variables[strings.Replace(arguments[1], "$", "", 1)])
			}
			if error1 != nil || error2 != nil {
				fmt.Println("ayak: failed to convert string to integer!")
				return
			}
			variables[strings.Replace(arguments[0], "$", "", 1)] = strconv.Itoa(integer + add)
		case _hlt:
			return
		}
	}
}
