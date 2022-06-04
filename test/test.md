# go test

go test 工具扫描 *_test.go 文件来寻找特殊函数，并生成一个临时的main包来调用它们，然后编译并运行，并汇报结果，最后清空临时文件。

## 测试函数

每一个测试文件必须导入`testing`包。这些函数的函数签名下如：

```go
func TestName(t *testing.T) {
    // ...
}
```

功能测试函数必须以`Test`开头，可选的后缀名称必须以大写的字母开头。

`go test`命令在不指定包参数的情况下，以当前目录所在的包为参数。使用下面的命令来编译和测试：

```shell
$ cd path/to/go/project
$ go test 
```

使用`go test -v`可以输出包中每个测试用例的名称和执行的时间。`-run` 的参数是一个正则表达式，它可以使得`go test`只运行那些
测试函数名称匹配给定模式的函数：

```shell
$ cd path/to/go/project
$ go test -v -run="Foo|Bar"
```

