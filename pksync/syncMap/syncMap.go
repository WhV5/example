/**
* @Author : henry
* @Data: 2020-07-21 17:06
* @Note: sync.Map demo
**/

package main

import (
	"fmt"
	"sync"
)

func main() {
	var sMap sync.Map
	sMap.Store(18, "henry")          // Store  添加元素
	if v1, ok := sMap.Load(18); ok { // Load 获取value
		fmt.Println(v1)
	}

	if v2, ok := sMap.LoadOrStore(18, "Cindy"); ok {
		fmt.Println("Exists 18-", v2) // LoadOrStore 获取 或 保存
		// 存在则获取，不存在则保存
	} else {
		fmt.Println("No 18-Cindy")
	}

	if v2, ok := sMap.LoadOrStore(20, "Cindy"); ok {
		fmt.Println("Exists 20-", v2)
	} else {
		fmt.Println("No 20-Cindy")
	}
	sMap.Range(func(k, v interface{}) bool { // 遍历map
		fmt.Printf("%d : %s\n", k, v)
		return true
	})

	sMap.Delete(18)
	sMap.Range(func(k, v interface{}) bool {
		fmt.Printf("%d : %s\n", k, v)
		return true
	})
}
