# 一致性hash算法

## 实现
* [consistent-hashing.go](../consistent-hashing.go)
* [consistent-hashing_test.go](../consistent-hashing_test.go)
```golang
package algorithm

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type UInt32Slice []uint32

func (s UInt32Slice) Len() int {
	return len(s)
}

func (s UInt32Slice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s UInt32Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type Hash func(data []byte) uint32

type Map struct {
	hash     Hash              // 哈希算法，默认crc32
	replicas int               // 虚拟节点数
	keys     UInt32Slice       // 已排序的节点哈希切片
	hashMap  map[uint32]string // 节点key哈希(考虑了虚拟节点数) 和 节点key（例如节点IP) 的map
}

func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[uint32]string),
	}
	// 默认使用crc32算法
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

func (m *Map) IsEmpty() bool {
	return len(m.keys) == 0
}

// Add 方法用来添加缓存节点，参数为节点key，例如IP
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		// 计算所有虚拟节点的哈希值，并存入m.keys中，同时在m.hashMap中保存哈希值和key的map
		// m.keys 是为了排序然后在后续查找节点时候进行二分查找到虚拟节点
		// m.hashMap 是为了在定位虚拟节点后获取实际节点
		for i := 0; i < m.replicas; i++ {
			hash := m.hash([]byte(strconv.Itoa(i) + key))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	// 对所有虚拟节点的哈希值进行排序，方便之后进行二分查找到虚拟节点
	sort.Sort(m.keys)
}

// Get 方法根据给定的对象获取最靠近它的那个节点信息
func (m *Map) Get(key string) string {
	if m.IsEmpty() {
		return ""
	}

	hash := m.hash([]byte(key))

	// 通过二分查找获取最优节点，第一个节点哈希值大于对象哈希值的就是最优节点（虚拟节点）
	idx := sort.Search(len(m.keys), func(i int) bool { return m.keys[i] >= hash })

	// 如果查找结果大于节点哈希数组的最大索引，表示此时该对象哈希值位于最后一个节点之后，那么放入第一个节点中
	if idx == len(m.keys) {
		idx = 0
	}

	//返回实际节点
	return m.hashMap[m.keys[idx]]
}

```

## 参考
* [使用Go实现一致性哈希](https://www.jianshu.com/p/b26555301f8e)