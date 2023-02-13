package main

import "fmt"

//01.go实现while循环实现输出2 - 3 + 4 - 5 + 6 ... + 100 的和
//02.for循环实现9乘9乘法表
//03.golang实现冒泡排序
//04.golang实现快排
//05.判断101-200之间有多少个素数，并输出所有素数
//06.实现map的有序遍历
//07.字符串反转
//08.合并两个有序列表
func sum_01() {
	sum := 0
	for i := 2; i < 101; i++ {
		if i%2 == 0 {
			sum += i
		} else {
			sum -= i
		}

	}
	fmt.Println(sum)

}

//乘法口诀表
func chengfang_02() {
	for y := 1; y <= 9; y++ {
		for x := 1; x <= y; x++ {
			fmt.Printf("%d*%d=%d\t", x, y, x*y)
		}
		fmt.Println()
	}
}

//03.golang实现冒泡排序
func BubbleSort() {
	list := []int{1, 9, 10, 8, 7}

	n := len(list)

	//在第一轮没有交换
	didSwap := false

	//进行第N-1轮迭代
	for i := n - 1; i > 0; i-- {
		//每次从第一轮开始比较，比较到第i位就不比较了，因为前一轮已经排好序
		for j := 0; j < i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				didSwap = true
			}

		}
		//如果在第一轮中没有交换，说明已经拍好了序，直接返回
		if !didSwap {
			return
		}
	}
	fmt.Println(list)
}

//04.golang实现快排

func main() {
	BubbleSort()
}
