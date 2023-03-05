package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	s, sep := "", " "
	start_time := time.Now()

	for _, arg := range os.Args[1:] {
		s += sep + arg
	}
	fmt.Println(s)
	stop_time := time.Now()

	fmt.Print("Traditional method time: ")
	fmt.Println(stop_time.Sub(start_time))

	start_time = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	stop_time = time.Now()

	fmt.Print("Join method time: ")
	fmt.Println(stop_time.Sub(start_time))

}
