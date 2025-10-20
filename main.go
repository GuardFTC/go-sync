// Package main @Author:冯铁城 [17615007230@163.com] 2025-10-16 15:06:22
package main

import "go-sync/pattern"

func main() {

	//1.协程/channel/select测试
	//base.TestGoroutineAndChanAndSelect()

	//2.mutex锁测试
	//sync_test.TestMutex()

	//3.rwMutex锁测试
	//sync_test.TestRWMutex()

	//4.WaitGroup测试
	//sync_test.TestWaitGroup()

	//5.Once测试
	//sync_test.TestOnce()

	//6.Pool测试
	//sync_test.PoolTest()

	//7.AtomicAddInt测试
	//atomic_test.AddInt32Test()

	//8.AtomicCAS测试
	//atomic_test.TestCompareAndSwap()

	//9.AtomicLoadStore测试
	//atomic_test.LoadStoreTest()

	//10.并发模式 生产者-消费者模式测试
	//pattern.ProducerAndConsumerTest()

	//11.并发模式 工作池模式测试
	//pattern.WorkPoolTest()

	//12.并发模式 流水线模式测试
	pattern.PipelineTest()
}
