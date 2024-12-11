package main

import (
	"fmt"

	"github.com/rudra-coder-mc/learn-go-with-tdd/helloworld"
	"github.com/rudra-coder-mc/learn-go-with-tdd/integers"
	"github.com/rudra-coder-mc/learn-go-with-tdd/iteration"
)

func hello() {
	fmt.Println(helloworld.Hello("world", "spanish"))
}

func integer() {
	fmt.Println(integers.Add(2, 3))
}

func iterations() {
	fmt.Println(iteration.Repeat("atul", 3))
}

func main() {
	fmt.Println("hii ")

	hello()
	integer()
	iterations()
}
