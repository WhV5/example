/**
* @Author : henry
* @Data: 2020-07-20 17:07
* @Note: WaitGroup错误示例：negative WaitGroup counter
**/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1)
	}()

	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2)
	}()

	wg.Wait()
	fmt.Println("All goroutine complete...")

	time.Sleep(1 * time.Second) // 在这里停顿 1s 后，就会发现问题： negative WaitGroup counter
}
