package main

import (
	"fmt"
	"math/rand"
	"os"
)

func osExit(code int) {
	fmt.Println("this is not main function")
	os.Exit(code)
}

func main() {
	fmt.Println("this is main function")
	osExit(1)
	if 2 > 1 {
		for 1 > 0 {
			r := rand.Intn(10)
			switch r {
			case 5:
				fmt.Println("hit")
				os.Exit(1) // want "cannot call os.Exit in function main in package main"
			default:
				fmt.Println("miss")
			}
		}
	}
}
