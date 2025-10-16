// Package base @Author:冯铁城 [17615007230@163.com] 2025-10-16 16:26:53
package base

import (
	"log"
	"time"
)

// TestGoroutineAndChanAndSelect 测试goroutine和chan和select
func TestGoroutineAndChanAndSelect() {

	//1.创建数据通道
	dataCh := make(chan int, 1)

	//2.创建读取完成信号通道
	finishCh := make(chan bool, 1)

	//3.创建写操作协程，向通道写入数据
	go func() {
		dataCh <- 1
		log.Printf("write data to channel finish")
	}()

	//4.创建读操作协程，从通道读取数据
	go func() {
		i := <-dataCh
		log.Printf("read data from channel: %d", i)
		finishCh <- true
	}()

	//5.Select循环阻塞，超时时间为2s
	select {
	case <-finishCh:
		log.Printf("base test finish")
	case <-time.After(2 * time.Second):
		log.Printf("base test timeout")
	}
}
