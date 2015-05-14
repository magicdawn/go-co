package task

import "github.com/magicdawn/go-co"
import . "github.com/tj/go-debug"

// run with
// DEBUG=goco:demo:* go run map.go
var debug = Debug("goco:task:map")

//
// task.Map(array,func(item,index) *co.Task,concurrency)
//
func Map(
	items []interface{},
	fn func(interface{}, int) *co.Task,
	concurrency int) *co.Task {

	// control flow
	total := len(items)
	running := 0
	started := 0
	completed := 0

	// prepare taskRet
	taskRet := new(co.Task)
	taskRet.Channel = make(chan interface{}) // can't use `chan []interface{}`
	taskRet.Result = make([]interface{}, total)

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

				// new Task
				debug("starting %d", index)
				// return Task will not panic
				t := fn(item, index)

				// collect the result
				// do not panic in map function
				taskRet.Result.([]interface{})[index], _ = co.Await(t)

				// maintain control flow
				running--
				completed++

				// notify
				oncomplete()
			}(items[started], started)

			started++
			running++
		}
	}

	go oncomplete()

	// collect error here
	defer func() {
		if err := recover(); err != nil {
			taskRet.Error = err.(error)
		}
	}()

	return taskRet
}
