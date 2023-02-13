package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

//启动一个协程
//主线程中每个100毫秒打印一次，总共打印2次
//另外开启一个协程，打印10次
//情况一：打印是交替，证明是并行的
//情况二：开启的协程打印两次，就退出了（因为主线程退出了）
func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("Testing() 你好", i)
		time.Sleep(100 * time.Millisecond)

	}
	wg.Done()
}

func test2() {
	for i := 0; i < 2; i++ {
		fmt.Println("Test2() 你好")
		time.Sleep(100 * time.Millisecond)

	}
	wg.Done()
}

/*
	wg.Add(1)
	go test()
	wg.Add(1)
	go test2()
	wg.Wait()
	fmt.Printf("主线程退出-----------")
*/
//定义协程数

var hei sync.WaitGroup

func hello(i int) {
	defer hei.Done()
	fmt.Println("Hello Goroutine!", i)
}

/*
	for i := 0; i < 10; i++ {
		hei.Add(1)
		go hello(i)
	}
	hei.Wait()
*/

func main() {
	// 1、创建channel
	ch := make(chan int, 5)
	// 2、向channel放入数据
	ch <- 10
	ch <- 12
	close(ch)
	fmt.Println("发送成功", ch)
	// 3、向channel取值
	v1 := <-ch
	fmt.Println(v1)
	v2 := <-ch
	fmt.Println(v2)
	// 4、空channel取值报错
	v3 := <-ch
	fmt.Println("v3", v3)

	// 在某些场景下我们需要同时从多个通道接收数据,这个时候就可以用到golang中给我们提供的select多路复

	//1.定义一个管道 10个数据int
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	//2.定义一个管道 5个数据string
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}
	//使用select来获取channel里面的数据的时候不需要关闭channel
	for {
		select {
		case v := <-intChan:
			fmt.Printf("从 intChan 读取的数据%d\n", v)
			time.Sleep(time.Millisecond * 150)
		case v := <-stringChan:
			fmt.Printf("从 stringChan 读取的数据%v\n", v)
			time.Sleep(time.Millisecond * 150)
		default:
			fmt.Printf("数据获取完毕")
			return //注意退出...
		}
	}
}
