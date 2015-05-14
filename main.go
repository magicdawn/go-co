//
// async/await for glolang
//
package co

//
// Task definition
//
type Task struct {
	// communicate via Channel
	Channel chan interface{}

	// store the result of the Task
	Result interface{}
}

//
// create a new Task
//
// execute fn , save the result, send to channel
//
// e.g
// co.Async(func() interface{}{
// 	return val
// 	// val will be Task's Result
// })

func Async(fn func() interface{}) (t Task) {
	t.Channel = make(chan interface{})

	// run the task
	// collect the result
	// set as the ret Task's Channel
	go func() {
		t.Result = fn()
		t.Channel <- t.Result
	}()

	return t
}

// await a Task & return it's result
//
// e.g
// res := co.Await(Task)
//
func Await(t Task) interface{} {
	// when t.Channel is available
	// set result as await ret value
	// `result = await(Task)`
	t.Result = <-t.Channel
	return t.Result
}
