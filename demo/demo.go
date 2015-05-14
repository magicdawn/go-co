package main

import (
	"fmt"
	"github.com/magicdawn/go-co"
	"time"
)

//
// construct a Task manualy
//
func sleep(ms int64) *co.Task {
	t := new(co.Task)
	t.Channel = make(chan interface{}, 1)

	go func() {
		// sleep a while
		time.Sleep(time.Millisecond * time.Duration(ms))

		// task is done
		// 10 is the result
		// send via channel
		t.Result = nil
		t.Channel <- t.Result
	}()

	return t
}

//
// construct a Task via co.Async
// return val will be result
// just as tj's co()
//
func sleepAsync(ms int64) *co.Task {
	return co.Async(func() interface{} {
		time.Sleep(time.Millisecond * time.Duration(ms))
		return nil
	})
}

func somethingAsync() *co.Task {
	return co.Async(func() interface{} {
		fmt.Println("somethingAsync started,please wait ...")
		co.Await(sleep(2000))
		fmt.Println("somethingAsync done!")
		return 10
	})
}

func main() {
	fmt.Println("before sleep : ", time.Now())
	co.Await(sleep(1000))
	fmt.Println("after sleep : ", time.Now())

	fmt.Println("before sleepAsync : ", time.Now())
	co.Await(sleepAsync(2000))
	fmt.Println("after sleepAsync : ", time.Now())

	res, _ := co.Await(somethingAsync())
	fmt.Println(res)
}
