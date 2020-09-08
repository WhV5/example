/**
* @Author : henry
* @Data: 2020-07-20 17:55
* @Note: waitGroup示例
**/

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("hello world")
	}(&wg)

	wg.Wait()
}
