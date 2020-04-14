package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Dictionary of supported reserved words across several programming languages, this list can be added to by the user:
var triggers = []string{"void", "def", "function", "func", "int", "var", "double", "float", "string", "String", "bool", "boolean"}
var filepath string

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("target filepath: ")
	x, _ := reader.ReadString('\n')
	x = strings.TrimSpace(x)
	filepath = x
	commandLoop()
}

func commandLoop() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	term, _ := reader.ReadString('\n')
	term = strings.TrimSpace(term)

	if len(term) < 1 {
		commandLoop()
	}

	run(term)
	commandLoop()
}
