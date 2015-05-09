# go-co

## Goals

to simulate 
- yield/Promise in node.js, [co](https://github.com/tj/co)
- async/await/Task in C#,

* checkout blog http://magicdawn.ml/2015/01/10/generator-and-promise/ *


use like

```go
Co(func(){
	res := await(someTask)
})
```

---

其实我是来吐槽golang的...
呢吗静态语言没有泛型,导致下面的代码

一坨 interface{} 差评
```golang
func Co(
	fn func(func(Task) interface{}) interface{}, // any : fn(await)
) (t Task) {
```
