package main

import (
	"fmt"
	"errors"
)
/**
* 数据类型
*/


/*----------------Boolean---------------------*/
// 值为true或false
var flag bool = false
var enable, disable = true, false

func testfunc() {
	available := true
	fmt.Println("%v", available)
}


/*------------------数值类型--------------------*/
// 整形分有符号和无符号两种
/**
* 虽然也有int, uint但长度取决于运行环境
* 有符号
* rune, int8, int16, int32, int64, byte
* 无符号
* uint8, uint16, uint32, uint64
* rune是int32的简称,byte是int8的简称
*/

// Go不允许不同数据类型之间进行运算

// 浮点型
// 只有float32和float64,并没有float

// 复数类型
// 默认complex128,64位实数+64位虚数
// 也有小一点的complex64,32位实数+32位虚数
var cplx1 complex128 = 1 + 2i
func test1() {
	fmt.Println("cplx1 is :", cplx1) // cplx1 is : (1+2i)
}


//------------------------字符串类型---------------------//
var str string
// 字符串类型不可变
// 可以用双引号和反引号表示
func testStr() {
	str = "abc"
	// str[0] = "d" // cannot assign to str[0]
	// 那么如何获取修改某个下标字符的字符串呢
	bs := []byte(str)
	bs[0] = 'd'
	strM := string(bs)
	fmt.Println(strM) // dbc

	// go允许字符串之间运算
	addStr := "hello" + " " + "world"
	fmt.Println(addStr) // => hello world

	// 字符串可以用切片
	fmt.Println("c" + str[1:]) // => cbc

	// 可以使用反引号声明多行字符串
	multiLineStr := `我爱学习,
我在学习Golang.`
	fmt.Println(multiLineStr)
}


//----------------------error类型----------------------------//
func testErr() {
	err := errors.New("some error infos...")
	if err != nil {
		fmt.Println(err)
	}
}


//-----------------------一些技巧----------------------------//

// 分组声明技巧
/**
同时声明多的变量,常量,或导入包的时候
import (
	"fmt"
	"errors"
)

var (
	i int
	pi float64
	prefix string
)

const (
	i = 100
	Pi float64 = 3.141593654
	prefix string = "producation"
)
*/

// itoa 枚举
const (
	spring = iota // 0
	summer = iota
	fall = iota
	winter = iota // 3
)

const ( // 新的const iota重置
	z1, z2, z3 = iota, iota, iota // 同一行值相同
	one = iota
)

const (
	s1 = 'a' // 每一行占一个枚举值
	b = iota // 1
)

func testMv() {
	fmt.Println(fall) // 2
	fmt.Println(z3) // 0
	fmt.Println(one) // 1
	fmt.Println(s1, b) // 97 1
}

//------------------------集合类---------------------------//
// 数组
var arr [5]int

/*
* 声明时可以写数组长度,不同的长度的数组代表不同的数据类型,长度是类型的一部分
* 如[3]renu 和 [6][renu]不是同一种类型
* 
*/
func testArr() {
	 arr1 := [7]int{1, 2, 3} // 可以声明的个数少于长度
	 arr1[1] = 22 
	 fmt.Println(arr1[0], arr1[1], arr1[6], len(arr1)) // => 1 22 0 默认赋值0
	 arr2 := [...]int{2, 3, 4} // 省略长度
	 fmt.Println(arr2[2])

	 // 多维数组
	 mArr := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{}} // [4]int可以省略
	 fmt.Println(mArr[0][1], mArr[1][1]) // 2 0

}

// slice 切片,形式上就是没有长度的数组 引用类型
var sli []int
func testSli(){
	sli := []byte{'a', 'b', 'c', 'd'} // 可以用别的数组或slice初始化
	fmt.Println(sli[0]) // 97
	var sli1 []byte = sli[0: 1]
	fmt.Println(sli1[0]) // 97

/*
slice其实是一个结构体,由3部分组成
1. 指向某个数组的某个下标地租的指针
2. 容量 capital
3. 指向的那个数组从slice开始位置指针到最后位置的长度
*/
	sli2 :=  []rune{'a', 'b', 'c', 'd', 'e', 'f'}
	fmt.Println(len(sli2))
	fmt.Println(cap(sli2))
	sli3 := sli2[1: 4 : 6] // 3参数切片最后一个参数指定容量,容量为6-1
	fmt.Println(cap(sli3))

	// 为什么slice要设计容量这种东西,slice有动态数组之说
	sli3 = append(sli3, 6, 7)
	fmt.Println(sli3[len(sli3) -1]) // 7
	fmt.Println(sli2[5]) // 7改变了原数组
}

// map 两种声明方式
var map1 map[int]string

func testMap(){
	//map2 = make(map[string]int)
	map1 = make(map[int]string)
	map1[1] = "abc"
	fmt.Println(map1[1])
	delete(map1, 1)
	// map也是引用类型
	map2 := map1
	map2[2] = "def"
	fmt.Println(map1[2])
}


//------------------------new 和 make-----------------------------//
/**
make用于内建类型 slice map channel make(T,arges) 返回T
原因是使用make初始化他们比较复杂,内部填充的是适当的值
new用于各种类型,new(T)返回*T,指向T的零值
*/

func main() {
	test1()
	testStr()
	testErr()
	testMv()
	testArr()
	testSli()
	testMap()
}

