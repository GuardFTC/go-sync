// Package sync_test @Author:冯铁城 [17615007230@163.com] 2025-10-17 10:34:10
package sync_test

import (
	"log"
	"sync"
)

// onceTest 测试Once
type onceTest struct {
	i int
	sync.Once
}

// inc 递增
func (o *onceTest) inc() {
	o.Do(func() {
		o.i++
	})
}

// TestOnce 测试Once
func TestOnce() {

	//1.初始化onceTest
	var once onceTest

	//2.并发10000次执行
	for i := 0; i < 10000; i++ {
		go once.inc()
	}

	//3.输出i
	log.Printf("once do i = %d\n", once.i)
}
