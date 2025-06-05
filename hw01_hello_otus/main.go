package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

const ReversedWord string = "Hello, OTUS!"

func main() {
	fmt.Println(reverse.String(ReversedWord))
}
