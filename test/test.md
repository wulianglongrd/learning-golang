# go test

go test 工具扫描 *_test.go 文件来寻找特殊函数，并生成一个临时的main包来调用它们，然后编译并运行，并汇报结果，最后清空临时文件。

## Test 函数

每一个测试文件必须导入`testing`包。这些函数的函数签名下如：

```go
func TestName(t *testing.T) {
    // ...
}
```

功能测试函数必须以`Test`开头，可选的后缀名称必须以`大写的字母`开头。参数`t`提供了汇报测试失败和日志记录的功能。

```go
func TestSin(t *testing.T) { /* ... */ }
func TestCos(t *testing.T) { /* ... */ }
func TestLog(t *testing.T) { /* ... */ }
```

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

### 测试日志

可以使用`t.Error`函数打印错误日志；使用`t.Errorf`来格式化错误信息，`t.Errorf`输出的失败的测试用例信息没有包函整个跟踪栈信息，
也不会导致程序宕机或者终止执行，我们可以使用`t.Fatal`或者`t.Fatalf`函数来终止测试。

```go
t.Error("some error info")
t.Errorf("%s = %q, want %q", foo, bar, want)
```

在测试的代码里面不会使用`log.Fatal`或者`os.Exit`，因为这两个调用会阻止跟踪的过程，这两个函数可以认为是main函数的特权。

### 覆盖率

测试指在发现bug，而不是证明其不存在。无论有多少测试都无法证明一个包是没有bug的。在最好的情下，它们增强了我们的信心，这些包可以在很多重要的场景下使用。

一个测试套件覆盖待测试包的比例称为测试的`覆盖率`。覆盖率无法直接通过数量来衡量，任何事情都是动态的，即使最微小的程序也无法精确的测试。

`语据覆盖率`是一种最简单且广泛使用的方法之一。语句覆盖率是指部份语句在一次执行中至少执行一次。

- 直接打印覆盖率信息

```shell
$ cd /path/to/go/project
$ go test -v -cover
```

- 启动覆盖率数据收集并查看

```shell
$ cd /path/to/go/project
$ go test -coverprofile=c.out # 启动覆盖率收集
$ go tools cover -html=c.out # 查看覆盖率数据
```

实现语句100%覆盖听上去很宏伟，但是在实际情况下这并不可行，也不会行之有效。因为语句得以执行并不意味着这是没有bug的，
拥有复杂表达式的语句块必须使用不同的输入执行多次来覆盖相关用命。

测试基本上是实用主义行为，在编写测试的代价和本可以通过测试避免的错误造成的代价之间进行平衡。



## Benchmark 函数

`基准测试`就是在一定的工作负载之下检测程序性能的一种方法。在go里面，`基准测试函数`与`测试函数`类似，但是前缀是`Benchmark`并且拥有一个
`*testing.B`的参数用来提供大多数和`*testing.T`相同的方法，额外增加了一些怀性能检测相关的方法。

```shell
$ cd /path/to/go/project
$ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/wulianglongrd/learning-golang/test
cpu: Intel(R) Core(TM) i5-5287U CPU @ 2.90GHz
BenchmarkFactorialForLoop-4     	23406238	        47.62 ns/op
BenchmarkFactorialRecursion-4   	 5234120	       228.0 ns/op
PASS
ok  	github.com/wulianglongrd/learning-golang/test	2.602s
```

标记`-bench`的参数指定了要运行的基准测试，它是一个匹配Benchmark函数名称的正则表达式，它的默认值不匹配任何函数。

基准测试名称的数字后缀4 (BenchmarkFactorialForLoop-4）表示`GOMAXPROC`的值，这个对于并发基准测试很重要。 

报告解析：基准测试用时2.602s，BenchmarkFactorialForLoop-4 调用 23406238 次的平均耗时 47.62 ns/op，BenchmarkFactorialRecursion-4 调用
5234120 次的平均耗时 228.0 ns/op

命令行使用`-benchmem`标记在报告中打印内存分配统计数据。

```shell
$ cd /path/to/go/project
$ go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/wulianglongrd/learning-golang/test
cpu: Intel(R) Core(TM) i5-5287U CPU @ 2.90GHz
BenchmarkFactorialForLoop-4     	24916185	        47.64 ns/op	       0 B/op	       0 allocs/op
BenchmarkFactorialRecursion-4   	 5275239	       227.3 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/wulianglongrd/learning-golang/test	2.677s
```

## Example 函数

被`go test`特殊对待的第三种函数就是示例函数，它们的名字以`Example`开关，即没有参数，也没有返回值。

```go
func ExampleAdd() {
    fmt.Println(Add(1, 2))
    fmt.Println(Add(2, 3))
    // output:
    // 3
    // 500
}
```

示例函数的三个目的：
- 作为文档，基于web的文档服务器`godoc`可以将示例函数和它所演示的函数或包相关联
- 可以通过`go test`运行可执行性测试。其前提是如上示例格式的输出验证，默认格式不满足默认是PASS
- 提供手动实验代码。 在 golang.org 上的 godoc 文档服务器使用 go playground来让用户实验代码


