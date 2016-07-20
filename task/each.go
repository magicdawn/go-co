package task

import "github.com/magicdawn/go-co"

// Each(items, func(item,index) *co.Task)
// iterate on array
func Each(
	items []interface{},
	fn func(interface{}, int) *co.Task) *co.Task {

	// prepare taskRet
	var taskRet = new(co.Task)
	taskRet.Channel = make(chan interface{}) // can't use `chan []interface{}`
	taskRet.Result = make([]interface{}, len(items))

	go func() {
		for index, item := range items {
			var t = fn(item, index)
			taskRet.Result.([]interface{})[index], _ = co.Await(t)
		}

		// send the  result
		taskRet.Channel <- taskRet.Result
	}()

	return taskRet
}
