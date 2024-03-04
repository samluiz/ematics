package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Welcome to Ematics!\n")
	fmt.Printf("Feel free to type in commands\n")
	Start(os.Stdin, os.Stdout)
}