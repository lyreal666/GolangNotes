package study
/**
* 变量
*/

/* 定义变量 */

// 使用var var varName type
var num int16;

// 定义多个变量
var a, b, c int32;

// 定义变量并初始化
var number int32 = 66;

// 定义多个变量并赋值
var n1, n2, n3 int = 2, 3, 4;

// 忽略类型声明，让编译器自动推断变量类型
var num1 = 5;

// 定义变量的最简洁方法
func test()  {
	//v := 888;
}

/**
* 使用:=替换var type简短声明变量只能在函数内部使用
* 一般用var 定义局部变量
* go对已经声明但没有使用的变量在编译阶段报错
*/

/* 常量 */

// 标准格式
const cvar int16 = 7

// 其它例子
const PI = 3.141592654
const prefix = "env=producation"

/**
* Go可以指定相当多位的小数位,指定为float 32会自动缩短为32bit,float64类似
*/

