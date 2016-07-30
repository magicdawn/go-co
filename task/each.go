package task

import "github.com/magicdawn/go-co"

// Each iterate a array
//
// example:
//  task.Each(arr, func(item, index, arr) { ... })
func Each(
	items []interface{},
	fn func(interface{}, int, []interface{}) *co.Task) *co.Task {
	return co.Async(func() interface{} {
		ret := make([]interface{}, len(items))
		for index, item := range items {
			var err error
			t := fn(item, index, items)
			ret[index], err = co.Await(t)
			if err != nil {
				panic(err)
			}
		}
		return ret
	})
}
