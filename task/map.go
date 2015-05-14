package task

import "github.com/magicdawn/go-co"
import . "github.com/tj/go-debug"

// run with
// DEBUG=goco:demo:* go run map.go
var debug = Debug("goco:demo:map")

//
// task.Map(array,func(item,index) Task,concurrency)
//
func Map(
	items []interface{},
	fn func(interface{}, int) co.Task,
	concurrency int) (taskRet co.Task) {

	// control flow
	total := len(items)
	running := 0
	started := 0
	completed := 0

	// prepare taskRet
	taskRet.Channel = make(chan interface{}) // can't use `chan []interface{}`
	taskRet.Result = make([]interface{}, total)
	// result := []interface{}{}

	// test chan
	// taskRet.Channel <- []int{1, 2, 3, 4}

	// concurrency
	if concurrency < 1 {
		concurrency = 1
	}

	var oncomplete func()

	// oncomplete callback
	oncomplete = func() {
		if completed >= total {
			taskRet.Channel <- taskRet.Result
			return
		}

		for started < total && running < concurrency {

			// start
			go func(item interface{}, index int) {
				debug("starting %d", index)

				// new Task
				t := fn(item, index)

				// collect the result
				taskRet.Result.([]interface{})[index] = co.Await(t)

				running--
				completed++

				oncomplete()
			}(items[started], started)

			started++
			running++
		}
	}

	go oncomplete()

	return taskRet
}
