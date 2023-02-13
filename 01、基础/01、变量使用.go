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

func Slice_list() {
	var a []string
	fmt.Println(a)        //[]
	fmt.Println(a == nil) //true

	var b = []int{}       //声明一个整型切片并初始化
	fmt.Println(b)        //[]
	fmt.Println(b == nil) //false

	var c = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(c)              //[false true]
	fmt.Println(c == nil)       //false

	// append()添加元素和切片扩容
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v len:%d cap:%d ptr:%p\n", numSlice, len(numSlice),
			cap(numSlice), numSlice)
	}
}

func map_list() {
	userInfo := map[string]string{
		"username": "root",
		"password": "123456",
		"name":     "heian_99",
	}
	fmt.Println(userInfo) // map[password:123456 username:IT 营小王子]
	v, ok := userInfo["username"]
	if ok {
		fmt.Println(v) // zhangsan
	} else {
		fmt.Println("map中没有此元素")
	}
	delete(userInfo, "password") //将 password从 map 中删除
	fmt.Println(userInfo)        // map[username:root]

	scoreMap := map[string]int{
		"zhangsan": 24,
		"lisi":     26,
		"wangwu":   24,
	}
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
}

// Pointer Go 语言中的指针操作非常简单，我们只需要记住两个符号：&（取地址）和 *（根据地址取值）
func Pointer() {
	var d = 10
	fmt.Printf("%d \n", &d)  // &a 指针地址 (824633761976)
	fmt.Printf("%d \n", *&d) // *&a 指针取值 (10)
	fmt.Printf("%T \n", &d)  // %T 指针类型 (*int )

	var a = 10
	var b = &a
	var c = *&a
	fmt.Println(a) // 10 a的值
	fmt.Println(b) // 0xc00001e060 a变量的内存地址
	fmt.Println(c) // 10 *内存地址 取值
}

func makeNew() {
	//var userinfo map[string]string
	//userinfo["username"] = "张三"
	//fmt.Println(userinfo)

	a := make([]int, 3, 10) // 切片长度为 1，预留空间长度为 10
	a = append(a, 1)
	fmt.Printf("%v--%T \n", a, a) // [0 0 0]--[]int     值----切片本身

	var b = new([]int)
	//b = b.append(b,2)                  // 返回的是内存指针，所以不能直接 append
	*b = append(*b, 3)          // 必须通过 * 指针取值，才能进行 append 添加
	fmt.Printf("%v--%T", *b, b) // &[]--*[]string  内存的指针---内存指针
}
func main() {
	makeNew()
}
