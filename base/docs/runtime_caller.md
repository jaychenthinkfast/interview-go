好的，以下是使用 `runtime.Caller` 和 `runtime.FuncForPC` 来查看函数调用次数的详细实现。这种方法通过在代码中插入统计逻辑，手动记录函数调用次数。

---

### 使用 `runtime.Caller` 和 `runtime.FuncForPC` 统计函数调用次数

`runtime.Caller` 可以获取调用栈信息，而 `runtime.FuncForPC` 可以将程序计数器（PC）转换为函数名。我们可以在程序中定义一个辅助函数来跟踪调用，并在全局变量中记录统计结果。

#### 示例代码：

[runtime_caller.go](../runtime_caller.go)

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	callCounts = make(map[string]int) // 存储函数调用次数
	mu         sync.Mutex             // 确保线程安全
)

// trackCall 记录调用该函数的上层函数的调用次数
func trackCall() {
	// 获取调用者的函数信息，1 表示调用 trackCall 的函数，调用层次
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return
	}
	fn := runtime.FuncForPC(pc).Name() // 获取函数名

	// 线程安全地增加计数
	mu.Lock()
	callCounts[fn]++
	mu.Unlock()
}

// 示例函数 foo
func foo() {
	trackCall()
	fmt.Println("foo called")
}

// 示例函数 bar
func bar() {
	trackCall()
	fmt.Println("bar called")
}

func main() {
	// 多次调用函数以测试
	for i := 0; i < 3; i++ {
		foo()
		bar()
	}

	// 打印统计结果
	mu.Lock()
	for fn, count := range callCounts {
		fmt.Printf("Function %s called %d times\n", fn, count)
	}
	mu.Unlock()
}
```

#### 输出：

```
foo called
bar called
foo called
bar called
foo called
bar called
Function main.foo called 3 times
Function main.bar called 3 times
```

---

### 实现原理

1. **`runtime.Caller(1)`**:

   - 参数 `1` 表示获取调用 `trackCall` 的函数的信息（即上层调用者）。
   - 返回值包括程序计数器（PC）、文件名、行号和成功标志。
2. **`runtime.FuncForPC(pc).Name()`**:

   - 将 PC 转换为具体的函数名（如 `main.foo`）。
3. **线程安全**:

   - 使用 `sync.Mutex` 保护 `callCounts` 字典，避免并发写入时的竞争条件。

---

### 使用步骤

1. 在需要统计调用次数的函数中插入 `trackCall()`。
2. 在程序结束或需要查看统计结果时，访问 `callCounts` 并打印。

---

### 优点与局限

- **优点**:
  - 可以精确统计特定函数的调用次数。
  - 不依赖外部工具，纯代码实现。
- **局限**:
  - 需要手动在目标函数中插入 `trackCall()`，对代码有一定侵入性。
  - 如果忘记插入或函数较多，维护成本较高。

---

如果你只需要统计某些特定函数的调用次数，这种方法非常直接有效。
