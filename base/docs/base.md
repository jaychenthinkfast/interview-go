# 应知应会
## 声明和使用变量
在 Go 中，常量名称通常以混合大小写字母或全部大写字母书写。

如果声明了变量但未使用，Go 会抛出错误

## 基本数据类型
rune 只是 int32 数据类型的别名。 它用于表示 Unicode 字符（或 Unicode 码位）。
```
rune := 'G'
println(rune)
//71
```
可以通过查看 [Go 源代码](https://golang.org/src/builtin/builtin.go) 来了解每种类型的范围。 了解每种类型的范围可帮助你选择正确的数据类型。

初始化字符串变量，你需要在双引号（"）中定义值。 单引号（'）用于单个字符


## 参考
[Go 入门](https://docs.microsoft.com/zh-cn/learn/modules/go-get-started/)