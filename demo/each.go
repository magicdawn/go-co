package main

import (
	"fmt"
	"github.com/magicdawn/go-co"
	"github.com/magicdawn/go-co/task"
	"time"
)

func main() {
	var items = []interface{}{1, 2, 3, 4, 5}

	// new Task
	var t = task.Each(items, func(item interface{}, index int) *co.Task {
		return co.Async(func() interface{} {
			var s = item.(int)
			time.Sleep(time.Second * time.Duration(s))
			fmt.Println("[%s] = %s", index, item)
			return s
		})
	})

	fmt.Println(co.Await(t))
}
