// Package pattern @Author:冯铁城 [17615007230@163.com] 2025-10-20 14:49:56
package pattern

import (
	"log"
)

// generateNum 生成数字
func generateNum(nums ...int) <-chan int {

	//1.创建任务通道
	out := make(chan int)

	//2.创建协程，异步向channel中写入数据
	go func() {

		//3.写入完成后关闭channel
		defer close(out)

		//4.写入数据
		for _, num := range nums {
			out <- num
		}
	}()

	//5.返回channel
	return out
}

// square 计算平方
func square(in <-chan int) <-chan int {

	//1.创建任务通道
	out := make(chan int)

	//2.创建协程，异步向channel中写入数据
	go func() {

		//3.写入完成后关闭channel
		defer close(out)

		//4.写入数据
		for num := range in {
			out <- num * num
		}
	}()

	//5.返回channel
	return out
}

// filterOdd 筛选奇数
func filterOdd(in <-chan int) <-chan int {

	//1.创建任务通道
	out := make(chan int)

	//2.创建协程，异步向channel中写入数据
	go func() {

		//3.写入完成后关闭channel
		defer close(out)

		//4.写入数据
		for num := range in {
			if num%2 != 0 {
				out <- num
			}
		}
	}()

	//5.返回channel
	return out
}

// PipelineTest 流水线模式测试
func PipelineTest() {

	//1.生成数字
	nums := generateNum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	//2.计算平方
	squares := square(nums)

	//3.筛选奇数
	oddSquares := filterOdd(squares)

	//4.打印结果
	for num := range oddSquares {
		log.Printf("num: %d\n", num)
	}
}
