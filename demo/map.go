package main

import (
	"github.com/magicdawn/go-co"
	"github.com/magicdawn/go-co/task"
	"time"
)

func sleep(sec int) co.Task {
	return co.Async(func() interface{} {
		time.Sleep(time.Second * time.Duration(sec))
		return nil
	})
}

func main() {
	items := []int{1, 2, 3, 4}

	// with concurrency 2
	t := task.Map(items.([]interface{}), func(item interface{}, index int) co.Task {
		return co.Async(func() interface{} {
			co.Await(sleep(item.(int)))
			return item
		})
	}, 2)
}
