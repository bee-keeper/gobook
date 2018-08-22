package main

import (
	"fmt"
	"os"
	"strings"
)

func WithJoin(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}

func WithoutJoin(args []string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Println(s)
}

func main() {
	if len(os.Args) > 1 {
		WithJoin(os.Args[1:])
		WithoutJoin(os.Args[1:])
	}
}
