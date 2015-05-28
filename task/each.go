package task

import "github.com/magicdawn/go-co"
import . "github.com/tj/go-debug"

// run with
// DEBUG=goco:demo:* go run map.go
var debug = Debug("goco:task:each")

//
// task.Each(array,func(item,index) *co.Task)
//
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
			var err error
			taskRet.Result.([]interface{})[index], err = co.Await(t)
		}

		// send the  result
		taskRet.Channel <- taskRet.Result
	}()

	return taskRet
}
