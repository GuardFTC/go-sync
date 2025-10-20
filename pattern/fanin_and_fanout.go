// Package pattern @Author:冯铁城 [17615007230@163.com] 2025-10-20 15:35:30
package pattern

import (
	"log"
	"sync"
)

// fanOutSquare 扇出-计算平方
func fanOutSquare(id int, in <-chan int, out chan<- int, wg *sync.WaitGroup) {

	//1.确保释放waitGroup标识位
	defer func() {
		wg.Done()
		log.Printf("fanOutSquare-%v写入结束", id)
	}()

	//2.从输入channel中读取数据，计算平方，并写入输出channel中
	for v := range in {
		out <- v * v
	}
}

// FanOutTest 扇出模式测试
func FanOutTest() {

	//1.创建waitGroup
	wg := sync.WaitGroup{}

	//2.创建输入、输出channel
	in := make(chan int)
	out := make(chan int)

	//3.启动3个协程，进行扇出计算
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go fanOutSquare(i, in, out, &wg)
	}

	//4.创建协程，异步向输入channel中写入数据
	go func() {
		defer func() {
			close(in)
			log.Printf("input channel closed")
		}()
		for i := 0; i < 10; i++ {
			in <- i
		}
	}()

	//5.创建协程，异步从输出channel中读取数据
	go func() {
		defer log.Printf("output channel closed")
		for v := range out {
			log.Printf("square: %d\n", v)
		}
	}()

	//6.等待waitGroup标识位
	wg.Wait()

	//7.关闭输出channel
	close(out)
}
