/*
A function can take zero or more arguments.

In this example, add takes two parameters of type int.

Notice that the type comes after the variable name.

(For more about why types look the way they do, see the article on Go's declaration syntax.)

*/

package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

/*


When function parameters share the same type, you can omit the type from all but the last.

In this example, x and y parameter share the same type, therefore type of x can be omitted.

*/

func addFunctionWithConsecutiveTypeParameters(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
