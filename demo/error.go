package main

import (
	"errors"
	"fmt"
	"github.com/magicdawn/go-co"
	. "github.com/tj/go-debug"
	"time"
)

var debug = Debug("goco:demo:error")

func sleep(sec int64) *co.Task {
	return co.Async(func() interface{} {
		time.Sleep(time.Second * time.Duration(sec))
		return nil
	})
}

func someWork() *co.Task {
	return co.Async(func() interface{} {
		co.Await(sleep(3))

		panic(errors.New("some error happened"))

		return "OK"
	})
}

func main() {
	t := co.Async(func() interface{} {
		res, err := co.Await(someWork())
		fmt.Println("1", res, err)

		if err != nil {
			panic(err)
		}

		return "OK"
	})

	res, err := co.Await(t)
	fmt.Println("2", res, err)
}
