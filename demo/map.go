package main

import (
	"fmt"
	"github.com/magicdawn/go-co"
	"github.com/magicdawn/go-co/task"
	"time"
)

func sleep(sec int) *co.Task {
	return co.Async(func() interface{} {
		time.Sleep(time.Second * time.Duration(sec))
		return nil
	})
}

func main() {
	items := []interface{}{1, 2, 3, 4}
	fmt.Println("before work : ", time.Now())

	// with concurrency 2
	t := task.Map(items, func(item interface{}, index int) *co.Task {
		return co.Async(func() interface{} {
			co.Await(sleep(item.(int)))
			return item.(int) * item.(int)
		})
	}, 2)

	res, _ := co.Await(t)
	fmt.Println("after work : ", time.Now())

	fmt.Println(res)
}

// if concurrency = 1
// max work lasts 4 secs , task t lasts 4secs
//
// if concurrency = 2
//	time	worker1							worker2
// 	0			task1								task2
//
//	1s		task3								task4
//
//	2s		task3								task4
//
// 	3s		task3								task4
//
// 	4s		task3								task4
//
// 	5s												task4-done
//
// total: 5secs
//
