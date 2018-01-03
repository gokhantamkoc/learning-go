/*

There are two ways to import packages

Using factored import:

import (
	"fmt"
	"math"
)

Using single package import:

import "fmt"
import "math"

It is a good practice to use the factored import statement.

*/

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.", math.Sqrt(7))
}
