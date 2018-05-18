package main

import (
	"fmt"
)

/*
* 面向对象
*/

/**
* 自定义类型
* 使用type重定义的类型都是自定义类型
* 所有自定义类型可以通过声明接收者的方式实现某个方法
*/

type color byte
type bs []bytes


//----------------method---------------------//
/**
* method is functiono with an implicit first argument
* called a recieved
*/

type Rectangle struct {
	width, height float32
}

func (r Rectangle) area() float32 { // r为copy传值
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c *Circle) area() float64{ // c为传指针
	const PI float64 = 3.141592654
	return PI * c.radius * c.radius // 使用时不需要解引用
}

func testMethod() {
	r := Rectangle{5.0, 4.0}
	fmt.Println(r.area())

	c := Circle{5.0}
	fmt.Println(c.area()) // 使用时不需要对c取地址
}


//------------------method的继承---------------------//
// 方法如同struct字段一样会被继承
type Human struct {
	name string age int
}

func (h Human) say(str string) {
	fmt.Println("Human:" ,str)
}

type Programer struct {
	Human // Go的oop是通过组合的方式实现的
	programLang string
}

func (p Programer) say(str) { // 方法的重写
	fmt.Println("Programer", str)
}

// 通过方法名大小写实现访问权限的设置

func main() {
	testMethod()
}