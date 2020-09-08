/**
* @Author : henry
* @Data: 2020-07-21 16:26
* @Note: RWMutex 详细demo
**/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var rwm sync.RWMutex
	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Printf("Try to lock for reading...[%d]\n", i)
			rwm.RLock()
			fmt.Printf("Locked for reading:[%d]\n", i)
			time.Sleep(2 * time.Second)
			fmt.Printf("Try to unlock for reading...[%d]\n", i)
			rwm.RUnlock()
			fmt.Printf("Unlocked for reading:[%d]\n", i)
		}(i)
	}
	//time.Sleep(100 * time.Millisecond)
	fmt.Println("Try to lock for writing...")
	rwm.Lock()
	fmt.Println("Locked for writing.")
	time.Sleep(100 * time.Millisecond)
}
