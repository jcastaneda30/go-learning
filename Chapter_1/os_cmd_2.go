package main

import (
	"fmt"
	"os"
)

func main_k() {
	s, sep := "", ""

	for index, arg := range os.Args[0:] {
		fmt.Println(index)
		s += sep + arg
		sep = ""
	}

	fmt.Print(s)
}
