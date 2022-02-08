package main

import (
	"fmt"
	"socpe/packageone"
)

var one string = "package level one" // is shadowed by the block level one

func main() {
	var one string = "block level one"

	fmt.Println(one) // will print "block level one"
	myFunc()         // will print "the number one"

	fmt.Println(packageone.PublicVar)
	packageone.Exported()
}

func myFunc() {
	var one = "the  number one"
	fmt.Println(one)
}
