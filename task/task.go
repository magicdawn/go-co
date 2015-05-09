package task

import . "github.com/magicdawn/go-co"

/*

parallel task

*/
func Parallel(tasks []Task, concurrency int32) (taskRet Task) {
	taskRet.Channel = make(chan []interface{})
	taskRet.Result = make([]interface{})

	total := len(tasks)
	result := make([]interface{}, total)

	running := 0
	started := 0
	completed := 0

	if concurrency < 1 {
		concurrency = 1
	}

	oncomplete := func() {
		if completed >= total {
			t.Channel <- t.Result
			return
		}

		for started < total && running < concurrency {
			t := tasks[started]
			started++
			running++

			go func(index int, t Task) {
				taskRet.Result[index] = Await(t)

				running--
				completed++

				oncomplete()
			}(running, t)
		}
	}

	go oncomplete()
}
