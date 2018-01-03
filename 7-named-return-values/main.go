/*

In Go Lang, return values may have names.

Return values which have names are treated as variables

These names should be used to document the meaning of the return values.

A return statement without arguments returns the named return values. This is known as a "naked" return.

Naked return statements should be used only in short functions, as with the example shown here. 

It is a best practice not to use "naked return" in complex functions.
*/

package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
	/*
	Result:
	7 10
	*/
}
