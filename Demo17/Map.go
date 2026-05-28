package main

import (
	"fmt"
	"sync"
)

// 定义安全Map
type SafeMap[K comparable, V any] struct {
	data  map[K]V
	mutex sync.RWMutex
}

// 创建
func NewMap[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		data: make(map[K]V),
	}
}

// 设置键值对
func (sm *SafeMap[K, V]) Set(key K, value V) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	sm.data[key] = value
}

// 获取键
func (sm *SafeMap[K, V]) Get(key K) (value V, ok bool) {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()
	value, ok = sm.data[key]
	return
}

// 获取所有键
func (sm *SafeMap[K, V]) Keys() []K {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()
	keys := make([]K, 0)
	for k, _ := range sm.data {
		keys = append(keys, k)
	}
	return keys
}

// 获取所有值
func (sm *SafeMap[K, V]) Values() []V {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()
	values := make([]V, 0)
	for _, v := range sm.data {
		values = append(values, v)
	}
	return values
}

// 获取map大小
func (sm *SafeMap[K, V]) Size() int {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()
	return len(sm.data)
}

// 清空
func (sm *SafeMap[K, V]) Clear() {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	sm.data = make(map[K]V)
}

// 删除键
func (sm *SafeMap[K, V]) Delete(key K) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	delete(sm.data, key)
}

func main() {
	m := NewMap[int, string]()
	m.Set(2, "12")
	m.Set(3, "34")
	fmt.Println(m.Keys())
	fmt.Println(m.Values())
	m.Set(2, "22")
	fmt.Println(m.Values())
}
