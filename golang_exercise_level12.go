package main

import (
	"./dog"
	"fmt"
)

type knine struct {
	name string
	age  int
}

func main() {
	/*
		Exercise 1
		1. Create a dog package.
		2. The dog package should have an exported func “Years” which takes human years and turns them into dog years (1 human year = 7 dog years).
		3. Document  your code with comments. Use this code in func main.
		4. Run your program and make sure it works
		5. Run a local server with godoc and look at your documentation.
	*/
	bingo := knine{
		name: "Bingo",
		age:  dog.Years(42),
	}
	fmt.Println(bingo)

	/*
		Exercise 2

	*/

	/*
		Exercise 3

	*/
}
