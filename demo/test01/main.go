package main

import (
	"fmt"

	"github.com/magicdawn/go-co"
)

func main() {
	n, _ := co.Await(hello())
	fmt.Println(n.(int))
}

func hello() *co.Task {
	return co.Async(func() interface{} {
		return 1
	})
}
