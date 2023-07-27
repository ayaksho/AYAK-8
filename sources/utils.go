package main

type word struct {
	function_type      int
	function_arguments []string
}

type lexer_out struct {
	words     []word
	functions map[string]int
}

const (
	_fun  = iota
	_mov  = iota
	_int  = iota
	_jmp  = iota
	_db   = iota
	_cdb  = iota
	_add  = iota
	_addx = iota
	_hlt  = iota
)

const (
	ax = "ax"
	bx = "bx"
	cx = "cx"
)

const (
	io = "0x10"
)
