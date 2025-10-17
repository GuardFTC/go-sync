// Package atomic_test @Author:冯铁城 [17615007230@163.com] 2025-10-17 14:17:04
package atomic_test

import (
	"log"
	"sync"
	"sync/atomic"
)

// AddInt32Test 测试AddInt
func AddInt32Test() {

	//1.创建WaitGroup
	var wg sync.WaitGroup

	//2.并发10000次执行,不使用原子类
	var resultWithOutAtomic int32
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resultWithOutAtomic++
		}()
	}

	//3.并发10000次执行,使用原子类
	var resultWithAtomic int32
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&resultWithAtomic, 1)
		}()
	}

	//4.等地执行完成
	wg.Wait()

	//5.输出最终结果
	log.Printf("不使用原子类 int32: %d", resultWithOutAtomic)
	log.Printf("使用原子类 int32: %d", resultWithAtomic)
}
