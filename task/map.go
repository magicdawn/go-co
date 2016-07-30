package task

import "github.com/magicdawn/go-co"

// Map : task.Map(array,func(item,index) *co.Task,concurrency)
func Map(
	items []interface{},
	fn func(interface{}, int, []interface{}) *co.Task,
	concurrency int) *co.Task {
	return co.Async(func() interface{} {
		if concurrency < 0 {
			concurrency = 1
		}

		// control flow
		total := len(items)
		running := 0
		started := 0
		completed := 0
		ret := make([]interface{}, total)
		chComplete := make(chan int, 1)
		chError := make(chan error, total)
		returned := false

		// oncomplete callback
		var oncomplete func()
		oncomplete = func() {
			if returned {
				return
			}

			if completed >= total {
				chComplete <- 1
				return
			}

			for started < total && running < concurrency {
				if returned {
					break
				}

				go func(item interface{}, index int) {
					// new Task
					debug("starting %d", index)

					var err error
					t := fn(item, index, items)
					ret[index], err = co.Await(t)
					if err != nil {
						chError <- err
						return
					}

					// notify
					running--
					completed++
					oncomplete()
				}(items[started], started)

				started++
				running++
			}
		}

		go oncomplete()

		select {
		case <-chComplete:
			returned = true
			return ret
		case err := <-chError:
			returned = true
			panic(err)
		}
	})
}
