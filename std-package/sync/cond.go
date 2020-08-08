package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	locker := new(sync.Mutex)
	cond := sync.NewCond(locker)

	for i := 0; i < 5; i++ {
		go func(i int) {
			cond.L.Lock()
			fmt.Println(i, " 获取锁")
			defer cond.L.Unlock()
			cond.Wait()
			fmt.Println(i, " 被唤醒")
			time.Sleep(time.Second)
		}(i)
	}

	fmt.Println("Signal......")
	cond.Signal()
	time.Sleep(time.Second)
	cond.Signal()
	time.Sleep(time.Second)
	fmt.Println("Broadcase......")
	cond.Broadcast()
	time.Sleep(time.Minute)
}
