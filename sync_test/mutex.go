// Package sync_test @Author:冯铁城 [17615007230@163.com] 2025-10-16 16:44:48
package sync_test

import (
	"log"
	"sync"
	"time"
)

// mutexTest 测试Mutex
type mutexTest struct {
	i          int
	sync.Mutex //读读阻塞、读写阻塞、写写阻塞
}

// IncWithLock 加锁递增
func (m *mutexTest) incWithLock() {

	//1.加锁
	m.Lock()
	defer m.Unlock()

	//2.i增加
	m.i++
}

// IncWithoutLock 不加锁递增
func (m *mutexTest) incWithoutLock() {
	m.i++
}

// TestMutex 测试Mutex
func TestMutex() {

	//1.创建MutexTest对象
	mt := mutexTest{}

	//2.创建10000个协程，对i进行递增(加锁)
	for i := 0; i < 10000; i++ {
		go mt.incWithLock()
	}

	//3.输出i的值
	time.Sleep(1 * time.Second)
	lockI := mt.i
	log.Printf("inc with mutex lock i = %d\n", lockI)

	//4.重置i
	mt.i = 0

	//5.创建10000个协程，对i进行递增(不加锁)
	for i := 0; i < 10000; i++ {
		go mt.incWithoutLock()
	}

	//6.输出i的值
	time.Sleep(1 * time.Second)
	NotLockI := mt.i
	log.Printf("inc without mutex lock i = %d\n", NotLockI)
}
