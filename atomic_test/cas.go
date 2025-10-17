// Package atomic_test @Author:冯铁城 [17615007230@163.com] 2025-10-17 14:51:16
package atomic_test

import (
	"log"
	"sync"
	"sync/atomic"
	"time"
)

// TestCompareAndSwap 测试CompareAndSwap
func TestCompareAndSwap() {

	//1.创建waitGroup
	var wg sync.WaitGroup

	//2.创建CAS标识位
	var cas int32

	//2.创建3个协程，对于变量进行CAS操作
	for i := 0; i < 10; i++ {

		//3.waitGroup标识位++
		wg.Add(1)

		//4.创建协程
		go func() {

			//5.确保最终释放waitGroup标识位
			defer wg.Done()

			//5.自旋进行CAS操作
			for {

				//6.如果CAS成功，Sleep规定时间，控制台打印，状态位重置，退出循环
				if atomic.CompareAndSwapInt32(&cas, 0, 1) {
					time.Sleep(time.Duration(i*100) * time.Millisecond)
					log.Printf("%v CAS成功", i)
					atomic.CompareAndSwapInt32(&cas, 1, 0)
					break
				}
			}
		}()
	}

	//7.等待所有协程执行完毕
	wg.Wait()
}
