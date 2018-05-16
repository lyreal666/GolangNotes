package main

import (
	"fmt"
)

/**
* 流程控制
* Go流程控制很人性化
*/

//---------------------条件语句-----------------------------//
/**
* 特点: 
* 不需要圆括号
* 条件语句内可以定义块作用域变量
*/

func computeValue() (value int) {
	value = 4
	return
}

func testCondition() {
	if value := computeValue(); value == 4 { // 条件语句内定义变量
		fmt.Println(value)
	}


	if age := 21; age > 0 && age < 18 { // age是块作用域
		fmt.Println("未成年")
	} else if testVar := 8; age >= 18 && age < 60 {
		fmt.Println("中年人", "testVar: ", testVar)
	} else {
		fmt.Println("老年人")
	}
	
	// fmt.Println(age) // 报错
}

//-------------------goto--------------------------//
func testGoto() {
	flag := 0
TAG:
	fmt.Println(flag)
	flag++
	if flag < 5 {
		goto TAG
	}
}


//-------------------For-----------------------//
// Golang最强大的控制逻辑
func testFor() {
	// 普通的三段式
	OUT: // 深层循环嵌套可以用goto,也可以用break,使用break标签时必须在使用前声明
	for x := 0; x < 1000; x++ { 
		for y := 0; y < 1000; y++ {
			for z := 0; z < 1000; z++ {
				if value := x + y + z; value == 666 {
					fmt.Println(x, y, z)
					break OUT
				} else {
					continue // 纯粹只是想说明有continue这个语法
				}
			}
		}
	}
	
	// for 省略三段式中1,3段就是while循环了
	index := 0
	for index < 5 {
		fmt.Println(index)
		index++;
	}

	// 范围for或者说遍历集合
	sli := []string{"a", "b", "c"}
	for index, value := range sli {
		fmt.Println(index, "is", value)
	}

	for key, value := range map[string]rune{"item1": 33, "item2": 66} {
		fmt.Println(key + "is: ", value)
	}

	for _, value := range [4]int{1, 2, 3, 4} { // 不需要的返回值全丢给_
		fmt.Println(value)
	}
}


//------------------------switch-----------------------------//
// 默认是不会往下继续执行的
// 通过fallthrough显式声明往下执行
func testSwitch() {
	// linux进程状态标志
	const (
		O = iota
		S = iota
		R = iota
		Z = iota
		T = iota
	)

	status := R
	switch status {
		case O:
			fmt.Println("运行中")
		case S:
			fmt.Println("休眠")
		case R:
			fmt.Println("准备运行")
			fallthrough
		case Z:
			fmt.Println("僵死进程")
		case T:
			fmt.Println("停止运行")
		default:
			fmt.Println("别的记不住了")
	}
}

func main() {
	testCondition()
	testGoto()
	testFor()
	testSwitch()
}