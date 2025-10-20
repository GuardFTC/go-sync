// Package pattern @Author:冯铁城 [17615007230@163.com] 2025-10-20 10:58:04
package pattern

import (
	"context"
	"log"
	"sync"
	"time"
)

// producer 生产者
func producer(id int, ch chan int, wg *sync.WaitGroup, ctx context.Context) {

	//1.确保最终释放waitGroup标识位
	defer func() {
		log.Printf("消息发送完毕 生产者-%v退出", id)
		wg.Done()
	}()

	//2.循环发送消息,通过select监听上下文退出信号
	for i := 1; i <= 3; i++ {
		select {
		case <-ctx.Done():
			log.Printf("收到监听信号，生产者-%v退出", id)
			return
		default:
			message := i*100 + i
			ch <- message
			log.Printf("生产者-%v发送消息-%v", id, message)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// consumer 消费者
func consumer(id int, ch chan int, wg *sync.WaitGroup, ctx context.Context) {

	//1.确保最终释放waitGroup标识位
	defer wg.Done()

	//2.循环接收消息,通过select监听上下文退出信号
	for {
		select {
		case <-ctx.Done():
			log.Printf("收到监听信号，消费者-%v退出", id)
			return
		case message, ok := <-ch:
			if !ok {
				log.Printf("消费者-%v接收消息失败", id)
				return
			}
			log.Printf("消费者-%v接收消息-%v", id, message)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// ProducerAndConsumerTest 测试生产者消费者模式
func ProducerAndConsumerTest() {

	//1.创建通道
	ch := make(chan int, 3)

	//2.创建带取消的上下文
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//3.创建消费者waitGroup
	var consumerWg sync.WaitGroup

	//4.创建2个消费者
	for i := 0; i < 2; i++ {
		consumerWg.Add(1)
		go consumer(i, ch, &consumerWg, ctx)
	}

	//5.创建生产者waitGroup
	var producerWg sync.WaitGroup

	//6.创建2个生产者
	for i := 0; i < 2; i++ {
		producerWg.Add(1)
		go producer(i, ch, &producerWg, ctx)
	}

	//6.等待生产者与消费者完成工作
	producerWg.Wait()

	//7.上下文取消，确保所有资源被释放
	cancel()

	//8.关闭通道
	close(ch)

	//9.等待所有消费者完成工作
	consumerWg.Wait()
}
