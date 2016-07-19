package main

import (
	"fmt"
	"time"

	"github.com/magicdawn/go-co"
)

func sleepAsync(ms int64) *co.Task {
	return co.Async(func() interface{} {
		time.Sleep(time.Millisecond * time.Duration(ms))
		return nil
	})
}

func main() {

	fmt.Println("now : ", time.Now())

	t := sleepAsync(2000).Continue(func(t *co.Task) interface{} {
		return 10
	})

	fmt.Println(co.Await(t))

	fmt.Println("now : ", time.Now())
}
