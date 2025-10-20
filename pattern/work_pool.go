// Package pattern @Author:冯铁城 [17615007230@163.com] 2025-10-20 11:59:53
package pattern

import (
	"fmt"
	"sync"
)

// Worker 工作者
type Worker struct {
	id      int
	tasksCh <-chan int
}

// work 工作
func (w *Worker) work(wg *sync.WaitGroup) {

	//1.最终释放waitGroup标识位
	defer wg.Done()

	//2.打印日志，声明已启动
	fmt.Printf("worker %d is started\n", w.id)

	//3.通过range从任务通道中获取任务，打印任务相关属性，直到通道关闭
	for task := range w.tasksCh {
		fmt.Printf("worker %d is working on task %d\n", w.id, task)
	}

	//4.打印日志，声明已结束
	fmt.Printf("worker %d is finished\n", w.id)
}

// NewWorker 创建一个工作者
func NewWorker(id int, tasksCh <-chan int) *Worker {
	return &Worker{
		id:      id,
		tasksCh: tasksCh,
	}
}

// NewWorkPool 创建一个工作池
func NewWorkPool(numWorkers int, wg *sync.WaitGroup) chan<- int {

	//1.创建任务通道
	tasksCh := make(chan int, 10)

	//2.循环创建worker
	for i := 0; i < numWorkers; i++ {

		//3.创建worker
		worker := NewWorker(i, tasksCh)

		//4.标志位++
		wg.Add(1)

		//5.启动worker
		go worker.work(wg)
	}

	//6.返回任务通道
	return tasksCh
}

func WorkPoolTest() {

	//1.声明worker数量
	numWorkers := 3

	//2.创建waitGroup
	var wg sync.WaitGroup

	//3.创建workPool
	tasksCh := NewWorkPool(numWorkers, &wg)

	//4.模拟传入100个任务
	for i := 0; i < 100; i++ {
		tasksCh <- i
	}

	//5.关闭通道
	close(tasksCh)

	//6.等待所有任务完成
	wg.Wait()
}
