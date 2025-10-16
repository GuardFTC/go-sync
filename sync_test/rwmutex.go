// Package sync_test @Author:冯铁城 [17615007230@163.com] 2025-10-16 17:12:45
package sync_test

import (
	"log"
	"sync"
	"time"
)

// rwMutexTest 测试RWMutex
type rwMutexTest struct {
	i            int
	sync.RWMutex //读读不阻塞、读写阻塞、写写阻塞
}

// incWithLock 加锁递增
func (r *rwMutexTest) incWithLock() {

	//1.加写锁
	r.Lock()
	defer r.Unlock()

	//2.i++
	r.i++
}

// incWithoutLock 不加锁递增
func (r *rwMutexTest) incWithoutLock() {
	r.i++
}

// readWithLock 读锁读取
func (r *rwMutexTest) readWithLock() int {

	//1.加读锁
	r.RLock()
	defer r.RUnlock()

	//2.返回
	return r.i
}

func TestRWMutex() {

	//1.创建rwMutexTest对象
	mt := rwMutexTest{}

	//2.创建10000个协程，对i进行递增(加锁)
	for i := 0; i < 10000; i++ {
		go mt.incWithLock()
	}

	//3.输出i的值
	time.Sleep(1 * time.Second)
	lockI := mt.i
	log.Printf("inc with rwmutex lock i = %d\n", lockI)

	//4.重置i
	mt.i = 0

	//5.创建10000个协程，对i进行递增(不加锁)
	for i := 0; i < 10000; i++ {
		go mt.incWithoutLock()
	}

	//6.输出i的值
	time.Sleep(1 * time.Second)
	NotLockI := mt.i
	log.Printf("inc without rwmutex lock i = %d\n", NotLockI)

	//7.创建10000个协程，对i进行读取(加读锁)
	for i := 0; i < 10000; i++ {
		go mt.readWithLock()
	}

	//8.等待读取完成
	time.Sleep(1 * time.Second)
	log.Printf("read with rwmutex lock success\n")
}
