package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, arg := range os.Args[1:] {
		fmt.Printf("%d "+arg+"\n", idx)
	}
}
