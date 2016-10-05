package co

/**
 * Task definition
 */

type Task struct {
	// communicate via Channel
	Channel chan interface{}

	// store the result of the Task
	Result interface{}

	// store error
	Error error
}

/**
 * Continue a task
 *
 * example:
 * t.Continue(func(t Task){
 *  // t is previous Task
 * }).Continue(func(t Task){
 *
 * })
 */

func (t *Task) Continue(fn func(*Task) interface{}) *Task {
	return Async(func() interface{} {
		_, err := Await(t)
		if err != nil {
			panic(err) // bump up error
		}

		// t is original Task
		return fn(t)
	})
}
