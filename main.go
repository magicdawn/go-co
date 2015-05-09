package co

// import "time"
// import "fmt"

type Task struct {
	Result chan interface{}
}

/**
 * create a new Task
 *
 * e.g
 * 	co.Async(func() interface{}{
 * 		return val
 * 		// val will be Task's Result
 * 	})
 */
func Async(fn func() interface{}) (t Task) {
	t.Result = make(chan interface{})

	// run the task
	// collect the result
	// set as the ret Task's Result
	go func() {
		result := fn()
		t.Result <- result
	}()

	return t
}

/**
 * await a Task & return it's result
 *
 * e.g
 * res := co.Await(Task)
 *
 * if `generic type exists`
 *
 * 	func Await(t Task<T>) T{
 * 		result := <- t.Result // t.Result is (chan T)
 * 		return result
 * 	}
 */
func Await(t Task) interface{} {
	// when t.Result is available
	// set result as await ret value
	// `result = await(Task)`
	result := <-t.Result
	return result
}
