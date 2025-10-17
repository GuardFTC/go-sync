package sync_test

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

// CoffeeCup 咖啡杯
type CoffeeCup struct {
	id      int64
	isUse   bool
	content string
}

// Use 使用咖啡杯
func (c *CoffeeCup) Use(content string) {
	c.isUse = true
	c.content = content
}

// Clear 清空咖啡杯
func (c *CoffeeCup) Clear() {
	c.isUse = false
	c.content = ""
}

// NewCoffeeCup 创建咖啡杯
func NewCoffeeCup(id int64) *CoffeeCup {
	return &CoffeeCup{
		id:      id,
		isUse:   false,
		content: "",
	}
}

// PoolTest 测试对象池
func PoolTest() {

	//1.初始化随机数种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	//2.创建对象池，模拟咖啡杯架
	pool := &sync.Pool{
		New: func() interface{} {
			id := time.Now().UnixNano()
			cup := NewCoffeeCup(id)
			log.Printf("可用咖啡杯不足，创建咖啡杯：%p, ID: %d", cup, cup.id)
			return cup
		},
	}

	//3.创建10个协程，模拟10个店员
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			//4.模拟一些延迟，增加复用机会
			randomId := getRandomDelay(r, 10)
			time.Sleep(time.Millisecond * 100 * time.Duration(randomId))

			//5.获取杯子
			cup := pool.Get().(*CoffeeCup)
			log.Printf("店员 %d: 从对象池获取咖啡杯：%p, ID: %d", id, cup, cup.id)

			//6.使用杯子
			cup.Use("咖啡")
			log.Printf("店员 %d: 使用咖啡杯：%p，内容：%s", id, cup, cup.content)

			//7.模拟使用时间
			time.Sleep(time.Millisecond * 50)

			//8.清洗杯子
			cup.Clear()
			log.Printf("店员 %d: 清洗咖啡杯：%p", id, cup)

			//9.放回架子
			pool.Put(cup)
			log.Printf("店员 %d: 咖啡杯：%p, ID: %d 已放回对象池", id, cup, cup.id)
		}(i)
	}

	//10.等待所有协程执行完毕
	wg.Wait()
}

// getRandomDelay 获取1到max范围内的随机数
// 如果max为0，则返回1
func getRandomDelay(r *rand.Rand, max int) int {
	if max == 0 {
		return 1
	}
	return r.Intn(max) + 1
}
