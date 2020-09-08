/**
* @Author : henry
* @Data: 2020-07-20 17:32
* @Note: waitGroup正确示例
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
		time.Sleep(1 * time.Second)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(1 * time.Second)
	}()

	wg.Wait()
	fmt.Println("All goroutine complete...")

	time.Sleep(3 * time.Second)
}
