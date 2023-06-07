package main

import "mangosteen/cmd"

/* 别名写法导出
import (
	c "mangosteen/cmd"
)
c.RunServer()调用
*/

func main() {
	cmd.RunServer()
}
