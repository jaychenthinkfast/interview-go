package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	callCounts = make(map[string]int)
	mu         sync.Mutex
)

func trackCall() {
	// 获取调用者的函数信息
	pc, _, _, ok := runtime.Caller(1) // 1 表示调用者的调用者,调用的层次
	if !ok {
		return
	}
	fn := runtime.FuncForPC(pc).Name()

	// 线程安全地记录调用次数
	mu.Lock()
	callCounts[fn]++
	mu.Unlock()
}

func foo() {
	trackCall()
	fmt.Println("foo called")
}

func bar() {
	trackCall()
	fmt.Println("bar called")
}

func main() {
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
