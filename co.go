// Package co : async/await for glolang
package co

// Async : create a new *Task
// execute fn , save the result, send to channel
//
// e.g
// co.Async(func() interface{}{
// 	return val
// 	// val will be Task's Result
// })
func Async(fn func() interface{}) *Task {
	t := new(Task)
	t.Channel = make(chan interface{})
	go func() {
		// final work
		defer func() {
			if err := recover(); err != nil {
				t.Error = err.(error)
			}

			// error is a finish state too
			t.Channel <- t.Result
		}()

		t.Result = fn()
	}()

	return t
}

// Await : await a Task & return it's result
//
// e.g
// res := co.Await(Task)
//
func Await(t *Task) (interface{}, error) {
	// when t.Channel is available
	// set result as await ret value
	t.Result = <-t.Channel

	return t.Result, t.Error
}
