/**
* @Author : henry
* @Data: 2020-07-20 21:51
* @Note: mutex demo 细节比较清晰
**/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var lock sync.Mutex
	fmt.Println("Lock the lock.(main)")
	lock.Lock()
	fmt.Println("The lock is locked.(main)")
	for i := 1; i <= 3; i++ {
		go func(i int) {
			fmt.Printf("Lock the lock.(g%d)\n", i)
			lock.Lock()
			defer lock.Unlock() // 加不加之后结果比较
			fmt.Printf("The lock is locked.(g%d)\n", i)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("Unlock the lock.(main)")
	lock.Unlock()
	fmt.Println("The lock is unlocked.(main)")
	time.Sleep(time.Second)
}
