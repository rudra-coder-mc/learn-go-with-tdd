package main

import (
	"fmt"

	"github.com/rudra-coder-mc/learn-go-with-tdd/arrays"
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
	fmt.Println(iteration.Repeat("atul ", 3))
}

func array() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(arrays.ArraySum(nums))
}

func main() {
	hello()
	integer()
	iterations()
	array()
}
