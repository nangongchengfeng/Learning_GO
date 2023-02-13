package main

import "fmt"

func main() {
	// 1.new实例化int
	age := new(int)
	*age = 1
	fmt.Println(*&age)
	// 2.new实例化切片
	li := new([]int)
	*li = append(*li, 1)
	fmt.Println(li)
	// 3.实例化map
	userinfo := new(map[string]string)
	*userinfo = map[string]string{}
	(*userinfo)["username"] = "张三"
	fmt.Println(userinfo) // &map[username:张三]
}
