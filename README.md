# go-co

[![GoDoc](https://godoc.org/github.com/magicdawn/go-co?status.svg)](https://godoc.org/github.com/magicdawn/go-co)

coroutine , async/await for golang

## Goals

to simulate
- yield/Promise in node.js, [co](https://github.com/tj/co)
- async/await/Task in C#

*checkout blog http://magicdawn.ml/2015/01/10/generator-and-promise/*

## API
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

*more checkout demo/ directory*

## more
其实我是来吐槽golang的...
- 呢吗静态语言没有泛型,导致一坨 interface{}
- 类型系统,这里本来定义了一个
    ```go
    type TaskDef struct {
        // blabla ...
    }

    type Task *TaskDef

    // 后面手动 new(TaskDef)应该就是*TaskDef = Task 类型
    // 不认
    ```
- 错误处理,defer panic, 常用的try/catch/throw 不用
    机制还是一样的，但是不好理解,defer recover相当于finally写在最前面
    如果已经panic,再defer 已经不起作用
- 差评!!!!!!!!!!!!!!!!!!

## License
the MIT license (magicdawn@qq.com)