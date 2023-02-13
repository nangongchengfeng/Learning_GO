package main

import "fmt"

//面的代码中定义了猫和狗，然后它们都会叫，你会发现main函数中明显有重复的代码
//如果我们后续再加上猪、青蛙等动物的话，我们的代码还会一直重复下去
//那我们能不能把它们当成“能叫的动物”来处理呢
type Cat struct {
	Name string
}

func (c Cat) Say() string {
	return c.Name + ": 喵喵喵"
}

type Dog struct {
	Name string
}

func (d Dog) Say() string { return d.Name + ": 汪汪汪" }
func Say() {
	c := Cat{Name: "小白猫"} // 小白猫：喵喵喵
	fmt.Println(c.Say())
	d := Dog{"阿黄"}
	fmt.Println(d.Say()) // 阿黄: 汪汪汪
}

//1.接口是一个规范
type Usber interface {
	start()
	stop()
}

//2.如果接口里面有方法的话，必要要通过结构体或者通过自定义类型实现这个接口
type Phone struct {
	Name string
}

//3.手机要实现usb接口的话必须得实现usb接口中的所有方法
func (p Phone) start() {
	fmt.Println(p.Name, "启动")
}
func (p Phone) stop() {
	fmt.Println(p.Name, "关机")
}

func Phone_acction() {
	p := Phone{
		Name: "华为手机",
	}
	var p1 Usber // golang中接口就是一个数据类型
	p1 = p       // 表示手机实现Usb接口
	p1.start()
	p1.stop()
}

//空接口作为函数的参数
func show(a interface{}) {
	fmt.Printf("值:%v 类型:%T\n", a, a)
}

func show_data() {
	show(20)         // 值:20 类型:int
	show("你好golang") // 值:你好golang 类型:string
	slice := []int{1, 2, 34, 4}
	show(slice) // 值:[1 2 34 4] 类型:[]int
}
func main() {
	show_data()
}
