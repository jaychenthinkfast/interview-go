# 查漏补缺
## 声明和使用变量
在 Go 中，常量名称通常以混合大小写字母或全部大写字母书写。

如果声明了变量但未使用，Go 会抛出错误。

## 基本数据类型
rune 只是 int32 数据类型的别名。 它用于表示 Unicode 字符（或 Unicode 码位）。
```
rune := 'G'
println(rune)
//71
```
可以通过查看 [Go 源代码](https://golang.org/src/builtin/builtin.go) 来了解每种类型的范围。 了解每种类型的范围可帮助你选择正确的数据类型。

初始化字符串变量，你需要在双引号（"）中定义值。 单引号（'）用于单个字符。

## 包
* 如需将某些内容设为专用内容，请以小写字母开始。
* 如需将某些内容设为公共内容，请以大写字母开始。

### 创建[模块](https://blog.golang.org/using-go-modules)
```
go mod init github.com/myuser/calculator
```
go.mod
```
module github.com/myuser/calculator

go 1.14
```

### 引用本地包（模块）
go.mod
``` 
module helloworld

go 1.14

require github.com/myuser/calculator v0.0.0

replace github.com/myuser/calculator => ../calculator
```
### 发布包
``` 
https://github.com/myuser/calculator

git tag v0.1.0
git push origin v0.1.0

import "github.com/myuser/calculator"
```

## switch 语句
### 省略条件
可以在 switch 语句中省略条件，就像在 if 语句中那样。 此模式类似于比较 true 值，就像强制 switch 语句一直运行一样。
``` 
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().Unix())
    r := rand.Float64()
    switch {
    case r > 0.1:
        fmt.Println("Common case, 90% of the time")
    default:
        fmt.Println("10% of the time")
    }
}
```

### 使逻辑进入到下一个 case
使用 fallthrough 关键字。
``` 
package main

import (
    "fmt"
)

func main() {
    switch num := 15; {
    case num < 50:
        fmt.Printf("%d is less than 50\n", num)
        fallthrough
    case num > 100:
        fmt.Printf("%d is greater than 100\n", num)
        fallthrough
    case num < 200:
        fmt.Printf("%d is less than 200", num)
    }
}
```
输出
```
15 is less than 50
15 is greater than 100
15 is less than 200
```
由于 num 为 15（小于 50），因此它与第一个 case 匹配。 但是，num 不大于 100。 由于第一个 case 语句包含 fallthrough 关键字，
因此逻辑会立即转到下一个 case 语句，**而不会对该 case 进行验证。 因此，在使用 fallthrough 关键字时必须谨慎。**

在某些编程语言中，你会在每个 case 语句末尾写一个 break 关键字。在 Go 中，当逻辑进入某个 case 时，它会退出 switch 块。

## panic 函数
内置 panic() 函数会停止正常的控制流。 所有推迟的函数（defer)调用都会正常运行。

## recover 函数
Go 提供内置函数 recover()，允许你在出现紧急状况之后重新获得控制权。 只能在已推迟的函数中使用此函数。 如果调用 recover() 函数，
则在正常运行的情况下，它会返回 nil，没有任何其他作用。

panic 和 recover 的组合是 Go 处理异常的惯用方式。 其他编程语言使用 try/catch 块。 Go 首选此处所述的方法。

提案阅读：[在 Go 中添加内置 try 函数的建议](https://go.googlesource.com/proposal/+/master/design/32437-try-builtin.md)

## 数组
如果你不知道你将需要多少个位置，但知道你将具有多少数据，那么还有一种声明和初始化数组的方法是使用省略号 (...)，如下例所示：
``` 
q := [...]int{1, 2, 3}
```

## 切片
append，Go 会自动扩充容量，**数量1024下时候扩容2倍，1024后扩容1.25倍。**

Go 具有内置函数 copy(dst, src []Type) 用于创建切片的副本。

## 映射
若要避免在将项添加到映射时出现问题，请确保使用 make 函数创建一个空映射（而不是 nil 映射）。
此规则仅适用于**添加项**的情况。 如果在 nil 映射中运行**查找、删除或循环操作**，Go 不会执行 panic。

## 方法
方法的一个关键方面在于，需要为任何类型定义方法，而不只是针对自定义类型（如结构）进行定义。 
但是，你不能通过属于其他包的类型来定义结构。 因此，不能在基本类型（如 string）上创建方法。
``` 
package main

import (
    "fmt"
    "strings"
)

type upperstring string

func (s upperstring) Upper() string {
    return strings.ToUpper(string(s))
}

func main() {
    s := upperstring("Learning Go!")
    fmt.Println(s)
    fmt.Println(s.Upper())
}
```
## 接口
Go 中的接口是一种抽象类型，只包括具体类型必须拥有或实现的**方法**。

## channel
如果希望 channel 仅发送数据，则必须在 channel 之后使用 <- 运算符。 如果希望 channel 接收数据，则必须在 channel 之前使用 <- 运算符
``` 
ch <- x // sends (or write) x through channel ch
x = <-ch // x receives (or reads) data sent to the channel ch
<-ch // receives data, but the result is discarded
```

关闭 channel 时，你希望数据将不再在该 channel 中发送。 **如果试图将数据发送到已关闭的 channel，则程序将发生严重错误**。
如果**试图通过已关闭的 channel 接收数据，则可以读取发送的所有数据。 随后的每次“读取”都将返回一个零值**。

### Channel 方向
``` 
chan<- int // it's a channel to only send data
<-chan int // it's a channel to only receive data
```
### 多路复用
select 语句的工作方式类似于 switch 语句，但它适用于 channel。 它会阻止程序的执行，直到它收到要处理的事件。 如果它收到多个事件，则会随机选择一个。
## 编写测试
创建测试文件时，该文件的名称必须以 _test.go 结尾。 你可以将你想用的任何内容用作文件名的前半部分，但典型做法是使用你要测试的文件的名称。

此外，要编写的每个测试都必须是以 Test 开头的函数。 然后，你通常为你编写的测试编写一个描述性名称，例如 TestDeposit。


## 参考
[Go 入门](https://docs.microsoft.com/zh-cn/learn/modules/go-get-started/)