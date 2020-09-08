/**
* @Author : henry
* @Data: 2020-07-23 19:18
* @Note: nil channel 的操作
**/

package main

func main() {
	var dataStream chan interface{}

	//dataStream := make(chan interface{})

	// 只读channel测试
	readStream(dataStream)

	// 只写channel测试
	writeStream(dataStream)
}

// var dataStream chan interface{}
// 只声明未实例化的只读channel
func readStream(readStream <-chan interface{}) {
	<-readStream // fatal error: all goroutines are asleep - deadlock!
	//readStream <- 0   // 编译错误
	//close(readStream) // 编译错误
}

// var dataStream chan interface{}
// 只声明未实例化的只写channel
func writeStream(writeStream chan<- interface{}) {
	//<- writeStream      // 编译错误
	writeStream <- 0   // fatal error: all goroutines are asleep - deadlock!
	close(writeStream) // panic: close of nil channel
}
