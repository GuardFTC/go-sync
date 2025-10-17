// Package atomic_test @Author:冯铁城 [17615007230@163.com] 2025-10-17 16:12:13
package atomic_test

import (
	"log"
	"sync"
	"sync/atomic"
	"time"
)

// LoadStoreTest 测试LoadStore
func LoadStoreTest() {

	//1.定义结果变量
	var result int32

	//2.定义写协程结束标识
	var isFinishWrite int32

	//3.定义读协程waitGroup
	var loadWg sync.WaitGroup

	//4.创建协程进行读操作
	loadWg.Add(1)
	go readValue(&loadWg, &isFinishWrite, &result)

	//5.定义写协程waitGroup
	var storeWg sync.WaitGroup

	//6.创建100个协程进行写操作
	for i := 0; i < 100; i++ {
		storeWg.Add(1)
		go func(i int) {
			defer storeWg.Done()
			time.Sleep(time.Millisecond * time.Duration(10*i))
			atomic.StoreInt32(&result, int32(i))
		}(i)
	}

	//7.等待写协程结束
	storeWg.Wait()
	log.Printf("write finish")

	//8.写协程结束标识符置为true
	atomic.StoreInt32(&isFinishWrite, 1)

	//9.等待读协程结束
	loadWg.Wait()
	log.Printf("read finish")
}

func readValue(loadWg *sync.WaitGroup, isFinishWrite *int32, result *int32) {

	//1.确保最终释放waitGroup标识位
	defer loadWg.Done()

	//2.定义暂存变量
	var loadTempValue int32

	//3.当写协程未结束时进行读操作
	for {

		//4.通过LoadInt32进行原子性读操作
		loadValue := atomic.LoadInt32(result)

		//5.如果读到的值和暂存变量不相等，则打印结果，暂存变量更新为当前值
		if loadValue != 0 && loadValue != loadTempValue {
			log.Printf("result is %d", loadValue)
			loadTempValue = loadValue
		}

		//6.如果写协程已结束，则跳出循环
		if atomic.LoadInt32(isFinishWrite) == 1 {
			break
		}
	}
}
