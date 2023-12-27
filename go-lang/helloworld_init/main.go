package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args

	if len(args) < 2 {
		fmt.Print("Usage has to pass 2 arguments after the hello run\n")
		os.Exit(1)
	}

	fmt.Printf("HelloWorld\nos.Args:%v\nArguments: %v\n", args, args[1:])

}
