package main

/*
This program is using the packages with import paths 'fmt' and 'math/rand'.
*/

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
