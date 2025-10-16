// Package sync_test @Author:冯铁城 [17615007230@163.com] 2025-10-16 17:38:15
package sync_test

import (
	"log"
	"sync"
)

// waitGroupTest 测试WaitGroup
type waitGroupTest struct {
	i int
	sync.WaitGroup
	sync.RWMutex
}

// inc 递增
func (w *waitGroupTest) inc() {

	//1.创建10000个协程，对i进行递增
	for i := 0; i < 10000; i++ {

		//2.主协程waitGroup+1
		w.Add(1)

		//3.创建协程，对i进行递增
		go func() {

			//4.加写锁
			w.Lock()
			defer w.Unlock()

			//5.i递增
			w.i++

			//6.释放waitGroup
			w.Done()
		}()
	}

	//7.阻塞式等待所有协程执行完毕
	w.Wait()
	log.Printf("wait group success i=%v\n", w.i)
}

// TestWaitGroup 测试WaitGroup
func TestWaitGroup() {
	wg := waitGroupTest{}
	wg.inc()
}
