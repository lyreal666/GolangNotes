package main 

import (
	"fmt"
)

/**
* 结构体struct
*/


// 定义结构体
type Student struct {
	name string
	age rune
	height float32
} 

func testStruct1() {
	// 普通声明
	var stu Student
	stu.name = "xiaoMing"
	stu.age = 21
	stu.height = 170.5

	// 不能遍历
	// for key, value := range stu {
	// 	fmt.Println(key, ": ", value)
	// }

	// 简短声明
	stu1 := Student{"xioaHong", 20, 160.2} // 值段声明
	stu2 := Student{name: "xiaoLan", age: 19, height: 168.6} // key-value方式声明

	fmt.Println("stu.name: ", stu.name)
	fmt.Println("stu1.age: ", stu1.age)
	fmt.Println("stu2.height: ",stu2.height)
	
	// 声明指针
	stuPtr := new(Student)
	(*stuPtr).name = "stuPointer" // *优先级没.高啊
	fmt.Println((*stuPtr).name)
}

//--------------------匿名字段----------------------//
// 类似js的属性拷贝
type Point struct {
	x int
	y int
}

type Circle struct {
	Point // 匿名字段
	r float32
}

func testAnonymous() {
	c := Circle{Point: Point{x: 4, y: 4}, r: 6.235}
	fmt.Println(c.r)
}

//---------------------字段重载---------------------//
type Header struct {
	H1 string
}

type H1 struct {
	span string
}

type H2 struct {
	span string
	small string
}

type Content struct {
	body string
}

type Dom struct {
	Header
	H1 H2 //Header里面的H1被外面的重载了
	Content
}

func testAcessFeild() {
	var dom Dom
	dom.H1 = H2{"<h1></h1>","<small></small>"}
	dom.Header.H1 = "<h1></h1>"
	fmt.Println(dom.H1) //{<h1></h1> <small></small>}
	fmt.Println(dom.Header.H1) // <h1></h1>
}

func main() {
	testStruct1()
	testAnonymous()
	testAcessFeild()
}