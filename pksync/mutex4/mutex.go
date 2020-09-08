/**
* @Author : henry
* @Data: 2020-07-21 15:56
* @Note:
**/

package main

import (
	"fmt"
	"sync"
	"time"
)

var j int

func main() {
	var lock sync.Mutex
	for i := 0; i < 5; i++ {
		go func(lock sync.Mutex) {
			lock.Lock()
			j++
			fmt.Printf("j = %d\n", j)
			lock.Unlock()
		}(lock)
	}
	time.Sleep(1 * time.Second)
}
