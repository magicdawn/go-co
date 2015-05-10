package co

// some extension on co.Task
//
// example:
//
// t
// .Continue(func(t Task){
//   t is previous Task
// })
// .Continue(func(t Task){
//
// })
func (t Task) Continue(
	fn func(Task) interface{}) Task {

	// return a wrapper Task
	return Async(func() interface{} {
		Await(t)

		// t is original Task
		return fn(t)
	})
}
