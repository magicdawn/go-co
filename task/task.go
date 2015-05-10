package task

import . "github.com/magicdawn/go-co"

/*
  task.Map(array,func(item,index) Task,concurrency)
*/
func Map(
	items []interface{},
	fn func(interface{}, int) Task,
	concurrency int) (taskRet Task) {

	// control flow
	total := len(items)
	running := 0
	started := 0
	completed := 0

	// prepare taskRet
	taskRet.Channel = make(chan interface{}, 1)
	taskRet.Result = make([]interface{}, total)

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
				t := fn(item, index)

				// collect the result
				taskRet.Result.([]interface{})[index] = Await(t)

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
