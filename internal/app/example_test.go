package app

import (
	"fmt"
	"math/rand"
)

func Example_makeURI() {
	uri := makeURI(rand.Intn(1000))
	fmt.Println(uri)
}
