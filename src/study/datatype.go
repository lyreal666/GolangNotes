package study

import "fmt"

/**
* 数据类型
*/

/*----------------Boolean---------------------*/
// 值为true或false
var flag bool = false
var enable, disable = true, false

func testfunc() {
	available := true
	fmt.Printf("%v", available)
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




