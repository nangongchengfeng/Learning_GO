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

//结构体
type person struct {
	name string
	city string
	age  int
}

func struct_name() {
	var p1 person
	p1.name = "张三"
	p1.city = "北京"
	p1.age = 18
	fmt.Printf("p1=%v\n", p1)  // p1={张三 北京 18}
	fmt.Printf("p1=%#v\n", p1) // p1=main.person{name:"张三", city:"北京", age:18}

	var p2 = new(person)
	p2.name = "张三"
	p2.age = 20
	p2.city = "北京"
	fmt.Printf("p2=%#v \n", p2) //p2=&main.person{name:"张三", city:"北京", age:20}

	p4 := person{
		name: "zhangsan",
		city: "北京",
		age:  18,
	}
	// p4=main.person{name:"zhangsan", city:"北京", age:18}
	fmt.Printf("p4=%#v\n", p4)
}

func struct_two() {
	name_p := Person{
		name: "heian",
		age:  18,
	}
	fmt.Println("变量：", name_p)
	name_p.printInfo()
	fmt.Printf("全部变量：num=%d\n", num) //num=10
}

type Person struct {
	name string
	age  int8
}

//值类型接受者
func (p Person) printInfo() {
	fmt.Printf("姓名:%v 年龄:%v\n", p.name, p.age) // 姓名:小王子 年龄:25
}

//指针类型接收者
func (p *Person) setInfo(name string, age int8) {
	p.name = name
	p.age = age
}

func main() {
	//struct_name()
	p1 := Person{
		name: "小王子",
		age:  25,
	}
	p1.printInfo() // 姓名:小王子 年龄:25
	p1.setInfo("张三", 20)
	p1.printInfo() // 姓名:张三 年龄:20
}
