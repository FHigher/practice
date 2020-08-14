package main

import (
	"fmt"
	"time"
)

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}
func main() {
	ch := make(chan int)
	go producer(ch)
	/* for {
		time.Sleep(1 * time.Second)
		v, ok := <-ch // 如果通道关闭，此处只是把ok值为false，但v拿到的是通道传递类型的零值，循环一直会运行下去

		fmt.Println("Received ", v, ok)
	} */

	// 通道关闭后，迭代结束，就停止了
	for v := range ch {
		time.Sleep(1 * time.Second)
		fmt.Println("Received", v)
	}
}
