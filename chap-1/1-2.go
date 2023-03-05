package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, arg := range os.Args[1:] {
		fmt.Print("Index: ")
		fmt.Print(idx)
		fmt.Println(" - " + "Arg: " + arg)
	}
}
