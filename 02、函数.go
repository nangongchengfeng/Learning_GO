package main

import (
	"fmt"
	"time"
)

//定义全局变量 num
var num int64 = 10

func init() {
	timeFormat1 := "2006-01-02 15:04:05"
	now := time.Now()
	fmt.Println("现在时间为：", now.Format(timeFormat1))
}
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

func sum_num() {
	plus, sub := calc(4, 5)
	fmt.Println(plus) // 和为：9
	fmt.Println(sub)  // 差为：-1
}
func main() {
	sum_num()
	fmt.Printf("全部变量：num=%d\n", num) //num=10
}
