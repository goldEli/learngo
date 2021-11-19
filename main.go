package main

import (
	"fmt"
	"strconv"
)

func queryUserById(id int) chan string {
	c := make(chan string)
	go func() {
		c <- "姓名" + strconv.Itoa(id)
	}()
	return c
}

func main() {
	c1, c2, c3 := queryUserById(1), queryUserById(2), queryUserById(3)
	c := make(chan string)
	// 开一个goroutine监视各个信道数据输出并收集数据到信道c
	go func() {
		for {
			// 监视c1, c2, c3的流出，并全部流入信道c
			select {
			case v1 := <-c1:
				c <- v1
			case v2 := <-c2:
				c <- v2
			case v3 := <-c3:
				c <- v3
			}
		}
	}()
	// 阻塞主线，取出信道c的数据
	for i := 0; i < 3; i++ {
		// 从打印来看我们的数据输出并不是严格的顺序
		fmt.Println(<-c)
	}
}
