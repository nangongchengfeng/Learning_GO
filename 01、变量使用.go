package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
	"unsafe"
)

//var 变量名 类型 = 表达式
var name string = "zhangsan"
var age int = 21

//var isOk bool bool

// 变量名 := 表达式
//批量声明变量

func ABC_string() {
	var (
		a string
		b int
		c bool
	)
	const (
		pi = 3.1415
		e  = 2.7182
	)
	const (
		n1 = 100
		n2
		n3
	)
	a = "张三"
	b = 10
	c = true
	fmt.Println(a, b, c, pi, e, n1, n2, n3)
}

func Printf_out() {
	fmt.Print("zhangsan", "lisi", "wangwu")   // zhangsanlisiwangwu
	fmt.Println("zhangsan", "lisi", "wangwu") // zhangsan lisi wangwu
	name := "zhangsan"
	age := 20
	fmt.Printf("%s 今年 %d 岁", name, age) // zhangsan 今年 20 岁
}

func init() {
	fmt.Println("这是初始化函数，系统自动调用")
}
func init() {
	timeFormat1 := "2006-01-02 15:04:05"
	now := time.Now()
	fmt.Println("现在时间为：", now.Format(timeFormat1))
}

func Byte_type() {
	var a int8 = 4
	var b int32 = 4
	var c int64 = 4
	d := 4
	fmt.Printf("a: %T %v \n", a, a)
	fmt.Printf("b: %T %v \n", b, b)
	fmt.Printf("c: %T %v \n", c, c)
	fmt.Printf("d: %T %v \n", d, d)
	fmt.Println(reflect.TypeOf(a))
	var e = true
	fmt.Println(e, "占用字节：", unsafe.Sizeof(b)) // true 占用字节： 1
}

func String_type() {
	s1 := "hello"
	s2 := "你好"
	s3 := `
  第一行
  第二行
  第三行`
	s := "美国第一"
	s_rune := []rune(s)
	fmt.Println(s, s_rune)
	fmt.Println("中国" + string(s_rune[2:])) // 中国第一
	fmt.Println(s1)
	fmt.Println(s1, s2, s3)
	s4 := "hello 张三"
	for i := 0; i < len(s4); i++ { //byte
		fmt.Printf("%v(%c) ", s4[i], s4[i])
		// 104(h) 101(e) 108(l) 108(l) 111(o) 32( ) 229(å) 188(¼) 160() 228(ä) 184(¸)

	}
	fmt.Println()          // 打印一个换行
	for _, r := range s4 { //rune
		fmt.Printf("%v=>%c ", r, r)
		// 104=>h 101=>e 108=>l 108=>l 111=>o 32=>  24352=>张 19977=>三
	}
	fmt.Println()

}

func to_string() {
	//1、int 转换成 string
	var num1 int = 20
	s1 := strconv.Itoa(num1)
	fmt.Printf("类型： %T ,值=%v \n", s1, s1) // 类型： string ,值=20
	// 2、float 转 string
	var num2 float64 = 20.113123
	/* 参数 1：要转换的值
	   参数 2：格式化类型
	   参数 3: 保留的小数点 -1（不对小数点格式化）
	   参数 4：格式化的类型
	*/
	s2 := strconv.FormatFloat(num2, 'f', 2, 64)
	fmt.Printf("类型： %T ,值=%v \n", s2, s2) // 类型： string ,值=20.11

	// 3、bool 转 string
	s3 := strconv.FormatBool(true)
	fmt.Printf("类型： %T ,值=%v \n", s3, s3) // 类型： string ,值=20.11
	//4、int64 转 string
	var num3 int64 = 20
	s4 := strconv.FormatInt(num3, 10)    /* 第二个参数10为 进制 */
	fmt.Printf("类型 %T ,值=%v \n", s4, s4) // 类型 string ,值=20

}

func string_to_init() {
	num := 100
	strNum := strconv.Itoa(num)
	fmt.Printf("num: %T %v \n", num, num)
	fmt.Printf("strNum: %T %v \n", strNum, strNum)

	intNum, _ := strconv.Atoi(strNum)
	fmt.Printf("intNum: %T %v \n", intNum, intNum)
}

func Array_list() {
	// 定义一个长度为 3 元素类型为 int 的数组 a
	var a [5]int
	// 定义一个长度为 3 元素类型为 int 的数组 b 并赋值
	var b [3]int

	b[0] = 80
	b[1] = 100
	b[2] = 96
	fmt.Println(a) // [0 0 0 0 0]
	fmt.Println(b) // [80 100 96]
	//普通遍历数组
	var c = [...]string{"北京", "上海", "深圳"}
	for i := 0; i < len(c); i++ {
		fmt.Println(c[i])
	}
	//k,v遍历数组

	var d = [...]string{"北京", "上海", "深圳"}
	for index, value := range d {
		fmt.Println(index, value)
	}
}
func main() {
	Array_list()
}
