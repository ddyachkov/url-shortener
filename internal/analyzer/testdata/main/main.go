package main

import (
	"fmt"
	"os"
)

func osExit(code int) {
	fmt.Println("this is not main function")
	os.Exit(code)
}

func main() {
	fmt.Println("this is main function")
	osExit(1)
	os.Exit(1) // want "cannot call os.Exit in function main in package main"
}
