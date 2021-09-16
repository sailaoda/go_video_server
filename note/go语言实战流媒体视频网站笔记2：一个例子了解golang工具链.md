---
title: go语言实战流媒体视频网站笔记2：一个例子了解golang工具链
cover:    /img/cover/goliu.jpg
tags:
  - - golang实战
categories:
  - - golang
    - go语言实战流媒体视频网站
abstract: '有东西被加密了, 请输入密码查看.'
message: '您好, 这里需要密码.'
wrong_pass_message: '抱歉, 这个密码看着不太对, 请再试试.'
wrong_hash_message: '抱歉, 这个文章不能被校验, 不过您还是能看看解密后的内容.'
date: 2021-09-15 10:47:16
keywords:
description: 通过一个简单的webservice具体从golang的工具链，到test，全面介绍golang在工程项目里需要掌握的知识点。
sticky:
password:
swiper_index:
---

# 一个Go的例子

{% note simple info %} 通过一个简单的webservice具体从golang的工具链，到test，全面介绍golang在工程项目里需要掌握的知识点。 {% endnote %}

## Go工具/命令

### Usage：

```go
go <command> [arguments]
```

### The commands are:

```go
        bug         start a bug report
        build       compile packages and dependencies
        clean       remove object files and cached files
        doc         show documentation for package or symbol
        env         print Go environment information
        fix         update packages to use new APIs
        fmt         gofmt (reformat) package sources
        generate    generate Go files by processing source
        get         add dependencies to current module and install them
        install     compile and install packages and dependencies
        list        list packages or modules
        mod         module maintenance
        run         compile and run Go program
        test        test packages
        tool        run specified go tool
        version     print Go version
        vet         report likely mistakes in packages
```

## 常用命令

```go
go build       compile packages and dependencies
```

- 最常用的go command之一，编译go文件。
- 跨平台编译：`env GOOS=linux GOARCH=amd64 go build`，编译目标平台的操作系统上运行的二进制文件。



```go
go install     compile and install packages and dependencies
```

- 也是编译，与build最大的区别是编译后会将输出文件打包成库放在pkg下

- 常用于本地打包编译的命令：go install



```go
go get        download and install packages and dependencies
```

- 用于获取go的第三方包，通常会默认从git repo 上pull 最新的版本
- 常用命令如： `go get -u github.com/go-sql-driver/mysql`（从github上获取mysql的driver并安装至本地，-u 表示下载最新的包。）



```go
go fmt         gofmt (reformat) package sources
```

- 类似于C中的lint，统一代码风格和排版
- 常用命令如：`go fmt`



```go
go test          test packages
```

- 运行当前包目录下的tests
- 常用命令如：`go test`或`go test -v`后面一个命令打印出所有信息



## Golang的test

- {% emp Go的test一般以XXX_test.go为文件名 %}

- XXX的部分一般为XXX_test.go所要测试的代码文件名。注：Go并没有要求XXX的部分必须是要测试的文件名。



文件main.go

```go
//main.go
package main
import(
	"io"
    "net/http"
)

func Printlto20() int{
    res := 0
    for i := 1 ;i <= 20 ; i ++{
        res += i
    }
    return res
}

func firstPage(w http.ResponseWriter,r *http.Request){
    io.WriteString(w, "<h1>Hello,this is my fisrt page!</h1>")
}

func main (){
    http.HandleFunc("/",firstPage)
    http.ListenAndServe(":8000",nil)
}


```

输出结果：

![程序输出结果](https://cdn.jsdelivr.net/gh/sailaoda/sai_img//img/3/image-20210915113343606.png)

文件测试：my_test.go（注意文件名格式 XXXtest.go）

```go
package main

import (
	"fmt"
	"testing"
)

func TestPrintlto20(t *testing.T) {
	res := Printlto20()
	fmt.Println("hey")
	if res != 210{
		t.Error("Wrong result of Printlto20")
	}
}
```

![测试结果](https://cdn.jsdelivr.net/gh/sailaoda/sai_img//img/3/image-20210915120054392.png)



# Test的写法

```go
package main
import (
	"testing"
)

func TestPrint(t *testing.T) {
    res := Printlto20()
    if res != 210 {
        t.Errorf("Return value not valid")
    }
}
```

- 每一个test文件必须import 一个testing，出现问题可以打印错误等
- test文件下的每一个test case均必须以Test开头并且符合TestXxx形式，否则go test会直接跳过测试不执行。  如 func TestPrint

### 具体细节

- test case 的入参为 t *testing.T 或 b *testing.B (b测试信息)
- t.Errorf为打印错误信息，并且当前test case 会被跳过
- t.SkipNow()为跳过当前test，并且直接按PASS处理下一个test。必须写在第一行。

第二点：改成200时的报错信息

```go
package main

import (
	"fmt"
	"testing"
)

func TestPrintlto20(t *testing.T) {
    //t.SkipNow()
	res := Printlto20()
	fmt.Println("hey")
	if res != 200{
		t.Error("Wrong result of Printlto20")
	}
}
```



![报错信息](https://cdn.jsdelivr.net/gh/sailaoda/sai_img//img/3/image-20210915131056792.png)

第三点：在第一行加上t.SkipNow()后，会跳过当前test，也会显示PASS

![Skip](https://cdn.jsdelivr.net/gh/sailaoda/sai_img//img/3/image-20210915131350235.png)

测试被忽略：1，已通过：0



## Test注意要点

- Go的test不会保证多个TestXxx是顺序执行，但是通常会按顺序执行。

- 使用t.Run 来执行subtests可以做到控制test输出以及test的顺序

例如

```go
package main 
import(
	"testing"
    "fmt"
)

func TestPrint(t *testing.T) {
    t.Run("a1",func(t *testing.T) {fmt.Println("a1")})
    t.Run("a2",func(t *testing.T) {fmt.Println("a2")})
    t.Run("a3",func(t *testing.T) {fmt.Println("a3")})
}
```



- 可能在做test之前要做一些初始化的东西
- 使用TestMain 作为初始化test，并且使用m.Run() 来调用其他tests可以完成一些需要初始化操作的testing，比如数据库连接，文件打开，REST服务登录等。

```go
func TestMain(m *testing.M) {
    fmt.Printkn("test main first")
    m.Run()
}
```

- {% emp 如果没有在TestMain中调用m.Run则除了TestMain意外的其他tests都不会被执行。 %}

例如：

### subtests测试实例

```go
package main 
import (
	"testing"
    "fmt"
)

func testPrint(t *testing.T) {                //一般情况下大小写区分开，作为子test
    res := Printlto20()
    fmt.Println("hey")
    if res != 210 {
        t.Errorf("Wrong result of Printlto20")
    }
}

func testPrint2(t *testing.T) {
    res := Printlto20()
    res ++
    if res != 211 {
        t.Errorf("Test Print2 failed")
    }
}

func TestAll(t *testing.T) {
    t.Run("TestPrint",testPrint)
    t.Run("TestPrint2",testPrint2)
}

```

测试结果：

```
=== RUN   TestPrint
hey
--- PASS: TestPrint (0.00s)
=== RUN   TestPrint2
--- PASS: TestPrint2 (0.00s)
=== RUN   TestAll
=== RUN   TestAll/TestPrint
hey
=== RUN   TestAll/TestPrint2
--- PASS: TestAll (0.00s)
    --- PASS: TestAll/TestPrint (0.00s)
    --- PASS: TestAll/TestPrint2 (0.00s)
PASS
ok  	webserver	5.658s

进程 已完成，退出代码为 0
```



子test小写后的输出结果：

```
=== RUN   TestAll
=== RUN   TestAll/TestPrint
hey
=== RUN   TestAll/TestPrint2
--- PASS: TestAll (0.00s)
    --- PASS: TestAll/TestPrint (0.00s)
    --- PASS: TestAll/TestPrint2 (0.00s)
PASS
ok  	webserver	5.576s

进程 已完成，退出代码为 0
```

### TestMain测试实例

```go
package main 
import (
	"testing"
    "fmt"
)

func testPrint(t *testing.T) {                //一般情况下大小写区分开，作为子test
    res := Printlto20()
    fmt.Println("hey")
    if res != 210 {
        t.Errorf("Wrong result of Printlto20")
    }
}

func testPrint2(t *testing.T) {
    res := Printlto20()
    res ++
    if res != 211 {
        t.Errorf("Test Print2 failed")
    }
}

func TestAll(t *testing.T) {
    t.Run("TestPrint",testPrint)
    t.Run("TestPrint2",testPrint2)
}

func TestMain(m *testing.M) {
    fmt.Println("Tests begins..... ")
    m.Run()
}
```

测试结果：

```
Tests begins..... 
=== RUN   TestPrintlto20
--- SKIP: TestPrintlto20 (0.00s)

测试已忽略.
=== RUN   TestAll
=== RUN   TestAll/TestPrint
hey
=== RUN   TestAll/TestPrint2
--- PASS: TestAll (0.00s)
    --- PASS: TestAll/TestPrint (0.00s)
    --- PASS: TestAll/TestPrint2 (0.00s)
PASS
ok  	webserver	4.981s

进程 已完成，退出代码为 0
```

如果把m.Run() 注释掉，则只会执行TestMain()函数的内容。



# Test之benchmark

- benchmark函数一般以Benchmark开头
- benchmark 的case 一般会跑b.N次，而且每次执行都会如此
- 在执行过程中会根据实际case的执行时间是否稳定会增加b.N的次数以达到稳态



```go
package main 
import (
	"testing"
    "fmt"
)

func testPrint(t *testing.T) {                //一般情况下大小写区分开，作为子test
    res := Printlto20()
    fmt.Println("hey")
    if res != 210 {
        t.Errorf("Wrong result of Printlto20")
    }
}

func testPrint2(t *testing.T) {
    res := Printlto20()
    res ++
    if res != 211 {
        t.Errorf("Test Print2 failed")
    }
}

func TestAll(t *testing.T) {
    t.Run("TestPrint",testPrint)
    t.Run("TestPrint2",testPrint2)
}

func TestMain(m *testing.M) {
    fmt.Println("Tests begins..... ")
    m.Run()
}

func BenchmarkAll(b *testing.B) {
    for n := 0 ; n < b.N ; n ++ {         //寻找函数最稳定的时候，得出最稳定的次数和每次跑的时间
        Printlto20()
    }
}
```

单独运行BenchmarkAll()  

`go test -bench=.`

```
Tests begins..... 
goos: windows
goarch: amd64
pkg: webserver
cpu: Intel(R) Core(TM) i7-1065G7 CPU @ 1.30GHz
BenchmarkAll
BenchmarkAll-8   	162938895	         7.531 ns/op
PASS
```

一共跑了162938895次，每次平均耗时7.531 ns/op。

### 注意：

1. Benchmark首字母大写

2. Benchmark同样也是普通的testcase之一，也会受TestMain限制。

3. 千万注意保证被测函数总能在一定时间达到稳态。

   





