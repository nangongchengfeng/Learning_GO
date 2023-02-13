package main

import (
	"fmt"
	"sync"
)

var x int64
var ng sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 15000; i++ {
		lock.Lock() // 加锁
		x = x + 1
		lock.Unlock() // 解锁
	}
	ng.Done()
}
func main() {
	ng.Add(3)
	go add()
	go add()
	go add()
	ng.Wait()
	fmt.Println(x) // 10000
}
