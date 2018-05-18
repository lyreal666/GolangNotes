package main

import (
	"fmt"
	"strconv"
)

/*
 Interface---一组方法的集合签名
*/

//----------------接口的定义------------------//
// 必须是普通方法的集合,传指针的方法不行
type Men interface {
	Say()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

type YoungChap interface {
	Say()
	Sing()
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	Say()
	Sing(song string)
	SpendSalary(amount float32)
}

type Human struct {
	name string
	age int
	phone string
}

func (human Human) Say() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", human.name, human.phone)
}

func (human Human) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}

func (human Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

type Student struct {
	Human //匿名字段Human
	school string
	loan float32
}

//Student实现BorrowMoney方法
func (s Student) BorrowMoney(amount float32) {
	s.loan += amount // (again and again and...)
}

type Employee struct {
	Human //匿名字段Human
	company string
	money float32
}

// Employee重载Human的Sayhi方法
func (e Employee) Say() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //此句可以分成多行
}

//Employee实现SpendSalary方法
func (e Employee) SpendSalary(amount float32) {
	e.money -= amount // More vodka please!!! Get me through the day!
}


//-------------------空接口---------------------------//
func testEmptyItfc() {
	var itfc interface{} // 空方法接口被任何类型实现
	ivar := 666 // 好像接口只能这样初始化
	str := "haha..."
	itfc = ivar
	itfc = str// 空方法接口被任何类型实现
	fmt.Println(itfc)
}

	
//------------------判断接口类型------------------------//
type Element interface{}
type List [] Element

type Person struct {
	name string
	age int
}

//定义了String方法，实现了fmt.Stringer
func (p Person) String() string {
	return "(name: " + p.name + " - age: "+strconv.Itoa(p.age)+ " years)"
}

func testCommaOk() {
	list := make(List, 3)
	list[0] = 1 // an int
	list[1] = "Hello" // a string
	list[2] = Person{"Dennis", 70}

	for index, element := range list {
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		} else {
			fmt.Printf("list[%d] is of a different type\n", index)
		}
	}
}

func testSwitchTest() {
	list := make(List, 3)
	list[0] = 1 // an int
	list[1] = "Hello" // a string
	list[2] = Person{"Dennis", 70}
	for index, element := range list{
		switch value := element.(type) {
			case int:
				fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
			case string:
				fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
			case Person:
				fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
			default:
				fmt.Println("list[%d] is of a different type", index)
		}
}

/*
 嵌入interface
Go里面真正吸引人的是它内置的逻辑语法，就像我们在学习Struct时学习的匿名字段，多么的优雅啊，那么相同的逻辑引入到interface里面，那不是更加完美了。如果一个interface1作为interface2的一个嵌入字段，那么interface2隐式的包含了interface1里面的method。

我们可以看到源码包container/heap里面有这样的一个定义

type Interface interface {
	sort.Interface //嵌入字段sort.Interface
	Push(x interface{}) //a Push method to push elements into the heap
	Pop() interface{} //a Pop elements that pops elements from the heap
}
我们看到sort.Interface其实就是嵌入字段，把sort.Interface的所有method给隐式的包含进来了。也就是下面三个方法：

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less returns whether the element with index i should sort
	// before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
另一个例子就是io包下面的 io.ReadWriter ，它包含了io包下面的Reader和Writer两个interface：

// io.ReadWriter
type ReadWriter interface {
	Reader
	Writer
}
反射
Go语言实现了反射，所谓反射就是能检查程序在运行时的状态。我们一般用到的包是reflect包。如何运用reflect包，官方的这篇文章详细的讲解了reflect包的实现原理，laws of reflection

使用reflect一般分成三步，下面简要的讲解一下：要去反射是一个类型的值(这些值都实现了空interface)，首先需要把它转化成reflect对象(reflect.Type或者reflect.Value，根据不同的情况调用不同的函数)。这两种获取方式如下：

t := reflect.TypeOf(i)    //得到类型的元数据,通过t我们能获取类型定义里面的所有元素
v := reflect.ValueOf(i)   //得到实际的值，通过v我们获取存储在里面的值，还可以去改变值
转化为reflect对象之后我们就可以进行一些操作了，也就是将reflect对象转化成相应的值，例如

tag := t.Elem().Field(0).Tag  //获取定义在struct里面的标签
name := v.Elem().Field(0).String()  //获取存储在第一个字段里面的值
获取反射值能返回相应的类型和数值

var x float64 = 3.4
v := reflect.ValueOf(x)
fmt.Println("type:", v.Type())
fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
fmt.Println("value:", v.Float())
最后，反射的话，那么反射的字段必须是可修改的，我们前面学习过传值和传引用，这个里面也是一样的道理。反射的字段必须是可读写的意思是，如果下面这样写，那么会发生错误

var x float64 = 3.4
v := reflect.ValueOf(x)
v.SetFloat(7.1)
如果要修改相应的值，必须这样写

var x float64 = 3.4
p := reflect.ValueOf(&x)
v := p.Elem()
v.SetFloat(7.1)
*/

func main() {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	//定义Men类型的变量i
	var i Men

	//i能存储Student
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.Say()
	i.Sing("November rain")

	//i也能存储Employee
	i = tom
	fmt.Println("This is tom, an Employee:")
	i.Say()
	i.Sing("Born to be wild")

	//定义了slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	//这三个都是不同类型的元素，但是他们实现了interface同一个接口
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x{
		value.Say()
	}

	testEmptyItfc()
	testCommaOk()
	testSwitchTest()
}
