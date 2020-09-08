/**
* @Author : henry
* @Data: 2020-07-26 21:37
* @Note: 生产者 消费者  也是一道面试题
**/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sendData(dataStream chan<- int) {
	defer close(dataStream)
	for i := 0; i < 5; i++ {
		dataStream <- rand.Intn(100)
	}
}

func recvData(dataStream <-chan int) {
	for data := range dataStream {
		time.Sleep(time.Second)
		fmt.Println(data)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var dataStream chan int
	dataStream = make(chan int)

	go func() {
		sendData(dataStream)
	}()

	recvData(dataStream)
}
