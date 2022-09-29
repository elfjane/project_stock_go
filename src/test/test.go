package main

import (
	"flag"
	"fmt"

	"elfjane.com/greetings"
)

func main() {
	filenamePtr := flag.String("ini", "sk88.ini", "a string")
	fmt.Println("Hello, World!2")
	greetings.Hello()
	fmt.Println(filenamePtr)
}
