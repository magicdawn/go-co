package main

import "time"
import "fmt"

type Task struct {
	Result chan interface{}
}

func sleep(ms int64) (t Task) {

	t.Result = make(chan interface{}, 1)

	go func() {
		// sleep 2 seconds
		time.Sleep(time.Millisecond * time.Duration(ms))

		// task is done
		// 10 is the result
		// send via channel
		t.Result <- 10
	}()

	return t
}

func Co(fn func(func(t Task) interface{})) (t Task) {

	await := func(t Task) interface{} {
		// when t.Result is available
		// set result as await ret value
		// `result = await(Task)`
		result := <-t.Result
		return result
	}

	// run the task
	fn(await)

	return t
}

func main() {
	Co(func(await func(Task) interface{}) {

		fmt.Println("before sleep : ", time.Now())

		res := await(sleep(1000))
		fmt.Println("result is ", res)

		fmt.Println("after sleep : ", time.Now())
	})
}
