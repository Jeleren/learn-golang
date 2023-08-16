/*
Go代码通过包来进行组织
每一个源文件都使用package声明
package 指明当前文件属于哪一个包
名为main的包是一个特殊的独立可执行的程序，而不是库
*/
package main

/*
在package声明后使用import导入需要的包
*/
import (
	"fmt"
	"osContent"
)

/*
main包中的main函数作为程序逻辑的开始
*/
func main() {
	fmt.Println("helloworld!")
	osContent.LogOs()
}