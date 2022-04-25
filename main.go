package main

import (
	_ "embed"
	"fmt"
)

//go:embed test.txt
var s string

func main() {
	fmt.Println("Hello, world!")
	fmt.Println(s)
}
