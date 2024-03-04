package main

import (
	"bufio"
	"fmt"
	"io"
)

const PROMPT = "> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	exit := false

	for !exit {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "exit" {
			exit = true
			continue
		}
		lexer := New(line)

		for token := lexer.ReadNextToken(); token.Type != EOF; token = lexer.ReadNextToken() {
			fmt.Printf("%+v\n", token)
		}
	}
}