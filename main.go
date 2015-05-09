package main

import "time"
import "fmt"

type Task struct {
	Result chan interface{}
}

func sleep(ms int64) (t Task) {

	t.Result = make(chan interface{}, 1)

	go func() {
		// sleep a while
		time.Sleep(time.Millisecond * time.Duration(ms))

		// task is done
		// 10 is the result
		// send via channel
		t.Result <- 10
	}()

	return t
}

func sleepAsync(ms int64) Task {
	return Co(func(await func(Task) interface{}) interface{} {
		time.Sleep(time.Millisecond * time.Duration(ms))
		return nil
	})
}

func Co(
	fn func(func(Task) interface{}) interface{}, // any : fn(await)
) (t Task) {

	t.Result = make(chan interface{})

	await := func(t Task) interface{} {
		// when t.Result is available
		// set result as await ret value
		// `result = await(Task)`
		result := <-t.Result
		return result
	}

	// run the task
	// collect the result
	// set as the ret Task's value
	go func() {
		result := fn(await)
		t.Result <- result
	}()

	return t
}

func main() {
	t := Co(func(await func(Task) interface{}) interface{} {

		fmt.Println("before sleep : ", time.Now())
		res := await(sleep(1000))
		fmt.Println("result is ", res)
		fmt.Println("after sleep : ", time.Now())

		fmt.Println("before sleepAsync : ", time.Now())
		res = await(sleepAsync(2000))
		fmt.Println("before sleepAsync : ", time.Now())

		return 1
	})

	<-t.Result
}
