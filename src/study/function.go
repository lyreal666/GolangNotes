package main

import (
	"fmt"
	"os"
)

/**
* 函数初步
*/

//------------------定义--------------------------//
// 声明返回值是一起声明变量名,可以直接在函数体使用
// 命名了返回值就不用显式用return语句声明了返回值了,但return不能省
// 也可以不命名返回值
// 类型声明时,如果有连续的同类型,可以只在最后一个变量后面声明,参数和返回值都适用
func normalFunc(a, b int, str1, str2 string) (retv1, retv2 int, retv3, retv4 string) {
	retv1, retv2, retv3, retv4 = a, b, str1, str2
	fmt.Printf("a: %d, b: %d, str1: %s, str2: %s\n", a, b, str1, str2)
	return // return 还是不能省略
}


//------------------多个返回值----------------------//
// 官方建议,如果函数是想导出的,就是给别人接口用,形式上是首字母大写,建议在声明时命名返回值
// 因为Golang不允许声明未使用的变量,所以如果返回值有不需要的可以用_占位
// Golang中一切赋值给_的值都会被舍弃
func sort(a, b rune) (smaller, bigger rune) {
	if a > b {
		smaller, bigger = b, a
	} else {
		smaller, bigger = a, b
	}
	return
}

//-------------------可变参数--------------------//
// 类似javascript的rest语法,可变参是slice类型
func restFunc(count int, args ...string) {
	outputStr := "传了" + string(count + '0') + "个字符"
	for _, arg := range args {
		outputStr += " " + arg
	}
	fmt.Println(outputStr)
}

//-------------------传值和传指针---------------------------//
// Go设计是为了取代c,c++的
// 所以Golang有指针,默认函数传参是值类型
// 使用传指针的方式可以修改实参
// 当参是数据比较大时还可以节省内存,毕竟指针一般就一个机器字长
func toUppercase(str *string) (upperStr string) {
	bytes := []byte(*str)

	for index := 0; index < len(bytes); index++ {
		if bytes[index] <= 'z' && bytes[index] >= 'a' {
			bytes[index] = 'A' + bytes[index] - 'a'
		} else {
			//
		}
	}
	upperStr = string(bytes)
	return
}

//---------------------defer-----------------------//
// 一种避免资源泄露的机制
// defer后面指定的语句在退出函数前执行
func testDefer() {
	fmt.Println("打开某个资源")
	defer fmt.Println("执行资源回收...")
	fmt.Println("利用资源在干活")
	fmt.Println("即将退出函数")
}

//------------------函数类型--------------------------//
type cmp func(int, int) (int)

func quickSort(arr []int, size int, cmpare cmp) {
	for index := 0; index < size - 1; index++ {
		minIndex := index
		for j := index + 1; j <= size -1; j++ {
			if cmpare(arr[index], arr[j]) > 0 {
				minIndex = j
			}
		}
		if minIndex != index {
			arr[index], arr[minIndex] = arr[minIndex], arr[index]
		}
	}
}

func cmpare(a, b int) (int) {
	return a - b
}


//-------------------异常处理------------------------//
// Go使用panic和recover进行异常处理
// 每个包的init函数在导入时初始化完所有常量和变量后会自动执行init函数
func init() {
	defer func() { // 如果panic后,直接退出函数前依次执行defer
		if e := recover(); e != nil {
			fmt.Println(e)			
			fmt.Println("接收并处理了异常")
		}
	}()
	user := os.Getenv("USER")
	if user == "" {
		panic("user is undefined!") // 类似抛异常
	}
}

func testRecover(f func()) (b bool) {
	defer func() {
		if e := recover(); e != nil { //使用recover相当于catch,返回值为捕获到的panic值
			b = true
		}
	}()
	f()
	return
}

func main() {
	normalFunc(666, 888, "chest", "butt")
	fmt.Println(sort(3, 2))
	restFunc(2, "hello", "world")
	str := "lowercase223333!"
	fmt.Println(toUppercase(&str))
	testDefer()
	arr := []int{3, 2, 6, 8, 9, 4, 20};
	
	quickSort(arr, 7, cmpare)
	for _, ele := range arr {
		fmt.Println(ele)
	}
}
