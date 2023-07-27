package main

import (
	"fmt"
	"strings"
)

func lex(lines []string) lexer_out {
	var out lexer_out
	out.functions = make(map[string]int)
	out.words = []word{}
	var line int

	for i := 0; i < len(lines); i++ {
		if lines[i] != "" {
			lines_raw := lines[i]

			for strings.Contains(lines_raw, "\t") {
				lines_raw = strings.Replace(lines_raw, "\t", "", 1)
			}
			for strings.Contains(lines_raw, "    ") {
				lines_raw = strings.Replace(lines_raw, "    ", "", 1)
			}

			words := strings.Split(lines_raw, " ")
			word_size := len(words[0])

			if words[0][word_size-1] == ':' {
				out.functions[strings.Replace(words[0], ":", "", 1)] = line

				var _word word
				_word.function_arguments = append(_word.function_arguments, strings.Replace(words[0], ":", "", 1))
				_word.function_type = _fun

				out.words = append(out.words, _word)
			}

			switch words[0] {
			case "mov":
				block := strings.Replace(words[1], ",", "", 1)
				data := words[2]

				var _word word
				_word.function_arguments = append(_word.function_arguments, block)
				_word.function_arguments = append(_word.function_arguments, data)
				_word.function_type = _mov

				out.words = append(out.words, _word)

			case "int":
				addres := words[1]

				var _word word
				_word.function_arguments = append(_word.function_arguments, addres)
				_word.function_type = _int

				out.words = append(out.words, _word)

			case "jmp":
				addres := strings.Replace(words[1], ",", "", 1)

				var _word word
				_word.function_arguments = append(_word.function_arguments, addres)
				_word.function_type = _jmp

				out.words = append(out.words, _word)

			case "db":
				name := strings.Replace(words[1], ",", "", 1)

				value := strings.Replace(lines_raw, fmt.Sprintf("db %s, ", name), "", 1)

				var _word word
				_word.function_arguments = append(_word.function_arguments, name)
				_word.function_arguments = append(_word.function_arguments, value)
				_word.function_type = _db

				out.words = append(out.words, _word)

			case "add":
				name := words[1]

				var _word word
				_word.function_arguments = append(_word.function_arguments, name)
				_word.function_type = _add

				out.words = append(out.words, _word)

			case "addx":
				name := strings.Replace(words[1], ",", "", 1)
				change_by := words[2]

				var _word word
				_word.function_arguments = append(_word.function_arguments, name)
				_word.function_arguments = append(_word.function_arguments, change_by)
				_word.function_type = _addx

				out.words = append(out.words, _word)

			case "hlt":
				var _word word
				_word.function_type = _hlt

				out.words = append(out.words, _word)
			}

			line = line + 1
		}
	}

	return out
}
