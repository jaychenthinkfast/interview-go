package algorithm

import (
	"testing"
)

//	go test -run=TestConsistentHashing -v
func TestConsistentHashing(t *testing.T) {
	//=== RUN   TestConsistentHashing
	//consistent-hashing_test.go:17: a 未检测到节点信息
	//consistent-hashing_test.go:20: a 192.168.0.2
	//consistent-hashing_test.go:21: b 192.168.0.1
	//--- PASS: TestConsistentHashing (0.00s)
	//PASS
	//ok      interview/algorithm     3.265s
	m := New(3, nil)
	if m.Get("a") == "" {
		t.Log("a", "未检测到节点信息")
	}
	m.Add("192.168.0.1", "192.168.0.2")
	t.Log("a", m.Get("a"))
	t.Log("b", m.Get("b"))
}
