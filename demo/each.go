package main

import (
	"fmt"
	"github.com/magicdawn/go-co"
	"github.com/magicdawn/go-co/task"
	. "github.com/tj/go-debug"
	"time"
)

func main() {
	var items = []int{1, 2, 3, 4, 5}

	// new Task
	var t = task.Each(items.([]interface{}), func(item interface{}, index int) *co.Task {
		return co.Async(func() interface{} {
			var s = item.(int)
			time.Sleep(time.Second * time.Duration(s))
			return s
		})
	})

	println(co.Await(t))
}
