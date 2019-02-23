package main

import (
	"fmt"
	
	"github.com/amelia678/mystringutils"
)

func main() {
	fmt.Println("Hello, World")
	s := "Hello, World"

	fmt.Println(mystringutils.Upper(s))
	fmt.Println(mystringutils.Lower(s))
}