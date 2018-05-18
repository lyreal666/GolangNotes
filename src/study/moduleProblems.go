package main

/**
* 关于import的知识
*/

// 分段式导入
// import "fmt"
// import "os"

// 分组声明
// import (
// 	"errors"
// 	"regexp"
// 	"io"
// )

// 导入模式 导入的时候注意接口名首字母大写
import (
	. "testPkg" // .表示使用接口时候不用写包前缀
	tpkg1 "testPkg1" // 给包定义别名
	_ "testPkg2" // 只执行init函数而不导入接口
)

func main() {
	TestImportPkg()
	tpkg1.TestImportMode()
	// testPkg2.testImportMode_() // undefined
}