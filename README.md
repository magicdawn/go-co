# go-co

[![GoDoc](https://godoc.org/github.com/magicdawn/go-co?status.svg)](https://godoc.org/github.com/magicdawn/go-co)

coroutine , async/await for golang

## Goals

to simulate
- yield/Promise in node.js, [co](https://github.com/tj/co)
- async/await/Task in C#

*checkout blog http://magicdawn.ml/2015/01/10/generator-and-promise/*

## Install
with [glide](https://github.com/Masterminds/glide)

```sh
$ glide get github.com/magicdawn/go-co
```

## API

### `co`
```go
import "github.com/magicdawn/go-co"
```

- co.Task : similar to Task in .NET,stands for a samll piece of work
- co.Await: await a Task

    ```go
    result,err := co.Await(task)
    ```
- co.Async: make a Task
    ```go
    func sleep() *co.Task{
    	return co.Async(func() interface{}{
        	time.sleep(time.Seconds * 10)
            return nil
        })
    }
    ```

    here is a Task will need 10 seconds

### `coutil`

- `coutil.Each` : like `Promise.each`
- `coutil.Map` : like `Promise.map` with concurrency support

*more checkout demo/ directory*

## CHANGELOG
[CHANGELOG.md](CHANGELOG.md)

## License
the MIT license (magicdawn@qq.com)